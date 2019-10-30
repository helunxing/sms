package processes

import (
	"encoding/json"
	"fmt"
	"net"
	"sms/common/message"
	"sms/common/utils"
)

// UserProcess 用户
type UserProcess struct {
	Conn net.Conn
}

// ServerProcessLogin 处理登陆请求
func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	tf := utils.Transfer{Conn: up.Conn}
	err = tf.WritePkg(data)
	fmt.Println("发送完毕，流程结束")
	return
}
