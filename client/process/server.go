package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sms/common/message"
	"sms/common/utils"
)

// ShowMenu 显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("-----用户xx登陆成功-----")
	fmt.Println("-----1.显示在线用户列表-----")
	fmt.Println("-----2.发送消息-----")
	fmt.Println("-----3.信息列表-----")
	fmt.Println("-----4.退出系统-----")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("在线用户")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入错误")
	}
}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	// 创建一个transfer实例，不停地读取服务器发送的消息
	tf := utils.Transfer{Conn: conn}

	for {
		fmt.Println("客户端正在等待读取读取")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		switch mes.Type {
		case message.NotifUserStatusMesType:
			// TODO 完成发送用户状态提醒
			var notifyUserStatusMes message.NotifUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// json.Unmarshal()
			updateUserStatus(&notifyUserStatusMes)
		default:
			fmt.Printf("mes类型未知，内容为 %s\n", mes)
		}
	}
}
