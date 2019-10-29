package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"sms/common/message"
	"sms/common/utils"
)

func main() {
	fmt.Println("服务器监听9999端口")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		fmt.Printf("与%s建立了连接\n", conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)
	}
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	// 声明resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 声明其内的LoginResMes
	var loginResMes message.LoginResMes

	// 判断用户合法性
	if loginMes.UserID == 100 && loginMes.UserPwd == "123" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 400
		loginResMes.Error = "400:用户名或密码错误"
	}

	// 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
		return
	}

	// 序列化后保存
	resMes.Data = string(data)

	// 再次序列化后发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail ", err)
		return
	}

	// 发送
	err = utils.WritePkg(conn, data)
	return
}

// serverProcessMes 根据客户端发送的消息种类不同，决定调用哪个函数处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		err = serverProcessLogin(conn, mes)
	default:
		err = errors.New("消息类型不存在，无法处理")
	}
	return
}

func process(conn net.Conn) (mes message.Message, err error) {
	defer conn.Close()
	for {
		mes, err = utils.ReadPkg(conn)
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Printf("自%s的客户端已断开连接\n", conn.RemoteAddr().String())
			default:
				fmt.Println("readPkg fail ", err)
			}
			return
		}
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println(err)
		}
	}
}
