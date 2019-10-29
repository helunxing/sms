package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
		fmt.Println("marshal err", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}
	// 发送消息
	err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("writepkg fail", err)
		return
	}
	// 接收消息
	mes, err = utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("readpkg fail", err)
		return
	}
	// 处理返回的数据
	var logResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &logResMes)
	if err != nil {
		fmt.Println("unmarshal fail", err)
		return
	}
	// 返回状态码
	if logResMes.Code != message.LoginResMesCodeOk {
		return errors.New(logResMes.Error)
	}
	return nil

}
