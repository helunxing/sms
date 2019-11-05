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
	Conn   net.Conn
	UserID int
}

// NotifyOthersOnlineUser 通知其他在线用户
func (up *UserProcess) NotifyOthersOnlineUser(userID int) {
	// 遍历 onlineUsers，逐个发送
	for id, onlineup := range userMgr.onlineUsers {
		if id == userID {
			continue
		}
		onlineup.NotifyMeOnline(userID)
	}
}

// NotifyMeOnline 通知其某用户已上线
func (up *UserProcess) NotifyMeOnline(userID int) {
	var mes message.Message
	mes.Type = message.NotifUserStatusMesType

	var notifUserStatusMes message.NotifUserStatusMes
	notifUserStatusMes.UserID = userID
	notifUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifUserStatusMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送失败", err)
		return
	}
}

// ServerProcessRegister 处理注册请求
func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	// 声明resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	// 声明其内的registerResMes
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ErrorUserExists {
			registerResMes.Code = 505
			registerResMes.Error = model.ErrorUserExists.Error()
		} else {
			registerResMes.Code = 505
			registerResMes.Error = "注册中未知错误"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.User.UserName, "注册成功")
	}

	// 将registerResMes序列化
	data, err := json.Marshal(registerResMes)
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
	fmt.Println("发送完毕，注册流程结束")
	return
}

// ServerProcessLogin 处理登陆请求
func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 取出并反序列化
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
			loginResMes.Code = message.LoginResMesCodeBadReq
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = message.LoginResMesCodeBadReq
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = message.LoginResMesCodeServerError
			loginResMes.Error = "未知内部错误"
		}
	} else {
		loginResMes.Code = 200
		up.UserID = loginMes.UserID
		userMgr.AddOnLineUser(up)
		up.NotifyOthersOnlineUser(loginMes.UserID)
		for id := range userMgr.onlineUsers {
			loginResMes.UsersID = append(loginResMes.UsersID, id)
		}
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
	fmt.Println("发送完毕，登陆流程结束")
	return
}
