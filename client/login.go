package main

import (
	"encoding/json"
	"errors"
	"sms/common/message"
	"sms/common/utils"
)

// 登陆校验
func login(userID int, userPwd string) (err error) {
	// 生成消息
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		return errors.New("marshal err " + err.Error())
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return errors.New("marshal err " + err.Error())
	}
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
	// 返回状态码
	if logResMes.Code != message.LoginResMesCodeOk {
		return errors.New(logResMes.Error)
	}
	return nil

}
