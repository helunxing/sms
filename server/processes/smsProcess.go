package processes

import (
	"encoding/json"
	"fmt"
	"net"
	"sms/common/message"
	"sms/common/utils"
)

// SmsProcess 处理消息
type SmsProcess struct {
}

// SendGroupMes 转发消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.unmarshal fail", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserID {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}
}

// SendMesToEachOnlineUser 向每个在线用户发消息
func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发失败", err)
	}
}
