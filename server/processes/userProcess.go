package processes

import (
	"encoding/json"
	"fmt"
	"net"
	"sms/common/message"
	"sms/common/utils"
	"sms/server/model"
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

	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNotExists {
			loginResMes.Code = 400
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = 401
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "未知内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
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
