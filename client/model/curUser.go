package model

import (
	"net"
	"sms/common/message"
)

// CurUser 客户端当前用户
type CurUser struct {
	Conn net.Conn
	message.User
}
