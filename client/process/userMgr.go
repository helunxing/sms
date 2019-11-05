package process

import (
	"fmt"
	"sms/common/message"
)

var onlineUsers = make(map[int]*message.User, 10)

// outputOnlineUser 显示当前用户列表
func outputOnlineUser() {
	for id := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// updateUserStatus 更新
func updateUserStatus(notifUserStatusMes *message.NotifUserStatusMes) {
	user, ok := onlineUsers[notifUserStatusMes.UserID]
	if ok {
		user.UserStatus = notifUserStatusMes.Status
	} else {
		onlineUsers[notifUserStatusMes.UserID] = &message.User{
			UserID:     notifUserStatusMes.UserID,
			UserStatus: notifUserStatusMes.Status,
		}
	}
}
