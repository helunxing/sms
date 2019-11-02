package message

// User 用户结构体
type User struct {
	UserID   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}
