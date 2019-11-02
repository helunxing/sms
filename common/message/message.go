package message

const (
	// LoginMesType 登陆消息Type字符串
	LoginMesType = "LoginMes"
	// LoginResMesType 登陆结果消息Type字符串
	LoginResMesType = "LoginResMes"
	// RegisterMesType 注册消息Type字符串
	RegisterMesType = "RegisterMes"
	// RegisterResMesType 注册结果消息Type字符串
	RegisterResMesType = "RegisterMesRes"
	// LoginResMesCodeOk 登陆成功消息码
	LoginResMesCodeOk = 200
	// LoginResMesCodeBadReq 请求错误消息码
	LoginResMesCodeBadReq = 400
	// LoginResMesCodeServerError 服务器错误消息码
	LoginResMesCodeServerError = 500
)

// Message 消息体
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// LoginMes 登陆消息
type LoginMes struct {
	UserID   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

// LoginResMes 登陆结果消息
type LoginResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// RegisterMes 注册消息
type RegisterMes struct {
	User User `json:"user"`
}

// RegisterResMes 注册响应
type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
