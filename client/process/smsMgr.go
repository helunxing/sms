package process

import (
	"encoding/json"
	"fmt"
	"sms/common/message"
)

func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return
	}
	info, err := fmt.Printf("用户id:\t%d 发送消息:\t%s", smsMes.UserID, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
