package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sms/common/message"
	"sms/common/utils"
	"sms/server/processes"
)

// Processor 连接处理器
type Processor struct {
	Conn net.Conn
}

func (p *Processor) process() (err error) {
	for {
		// 创建Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		var mes message.Message
		mes, err = tf.ReadPkg()
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Printf("自%s的客户端已断开连接\n", p.Conn.RemoteAddr().String())
				err = nil
			default:
				err = errors.New("readPkg fail " + err.Error())
			}
			return
		}
		err = p.serverProcessMes(&mes)
		return
	}
}

// serverProcessMes 根据客户端发送的消息种类不同，决定调用哪个函数处理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		up := processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		err = errors.New("消息类型不存在，无法处理")
	}
	return
}
