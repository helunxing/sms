package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"sms/common/message"
	"time"
)

// 登陆校验
func login(userID int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserID = userID
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	time.Sleep(10 * time.Second)
	fmt.Println("休眠了10秒")
	// TODO 处理返回的数据
	return
}
