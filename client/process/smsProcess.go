package process

import (
	"encoding/json"
	"fmt"
	"sms/common/message"
	"sms/common/utils"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserID = curUser.UserID
	smsMes.UserName = curUser.UserName
	smsMes.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err.Error())
		return
	}
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("send message err", err.Error())
	}
	return
}
