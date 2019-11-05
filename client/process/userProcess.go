package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sms/common/message"
	"sms/common/utils"
)

// UesrProcess 处理用户逻辑
type UesrProcess struct {
}

// Login 向服务器发送登陆
func (up *UesrProcess) Login(userID int, userPwd string) (err error) {

	// 生成消息
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd
	// 序列化登陆消息
	data, err := json.Marshal(loginMes)
	if err != nil {
		return errors.New("marshal err " + err.Error())
	}
	mes.Data = string(data)
	// 序列化消息
	data, err = json.Marshal(mes)
	if err != nil {
		return errors.New("marshal err " + err.Error())
	}
	// 打开连接
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	// 生成传输结构体
	tf := utils.Transfer{
		Conn: conn,
	}
	// 发送消息
	err = tf.WritePkg(data)
	if err != nil {
		return errors.New("writepkg fail " + err.Error())
	}
	// 接收消息
	mes, err = tf.ReadPkg()
	if err != nil {
		return errors.New("readpkg fail " + err.Error())
	}
	// 处理返回的数据
	var logResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &logResMes)
	if err != nil {
		return errors.New("unmarshal fail " + err.Error())
	}
	// 返回状态码信息
	if logResMes.Code != message.LoginResMesCodeOk {
		fmt.Printf("登陆失败，原因：%s\n", logResMes.Error)
		// fmt.Println("继续？y或Y继续，2注册，其他则退出：")
		// var chr string
		// fmt.Scanf("%d\n", &chr)
		// if !(chr == "y" || chr == "Y") {
		// 	break
		// }
		// if chr == "2" {
		// 	fmt.Println("注册逻辑")
		// }
	} else {
		fmt.Println("登陆成功")
		fmt.Println("当前在线用户列表如下：")
		for _, v := range logResMes.UsersID {
			if v == userID {
				continue
			}
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserID:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	}
	return
}

// Register 注册逻辑
func (up *UesrProcess) Register(userID int, userPwd string, userName string) (err error) {
	// 生成消息
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserID = userID
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	// 序列化登陆消息
	data, err := json.Marshal(registerMes)
	if err != nil {
		return errors.New("marshal err " + err.Error())
	}
	mes.Data = string(data)
	// 序列化消息
	data, err = json.Marshal(mes)
	if err != nil {
		return errors.New("marshal err " + err.Error())
	}
	// 打开连接
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	// 生成传输结构体
	tf := utils.Transfer{
		Conn: conn,
	}
	// 发送消息
	err = tf.WritePkg(data)
	if err != nil {
		return errors.New("writepkg fail " + err.Error())
	}
	// 接收消息
	mes, err = tf.ReadPkg()
	if err != nil {
		return errors.New("readpkg fail " + err.Error())
	}
	// 处理返回的数据
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		return errors.New("unmarshal fail " + err.Error())
	}
	// 返回状态码信息
	if registerResMes.Code != message.LoginResMesCodeOk {
		fmt.Printf("注册失败，原因：%s\n", registerResMes.Error)
	} else {
		fmt.Println("注册成功")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	}
	return
}
