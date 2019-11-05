package processes

import "fmt"

var (
	userMgr *UserMgr
)

// UserMgr 保存登陆的用户信息
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// AddOnLineUser 添加
func (um *UserMgr) AddOnLineUser(up *UserProcess) {
	um.onlineUsers[up.UserID] = up
}

// DelOnLineUser 删除
func (um *UserMgr) DelOnLineUser(UserID int) {
	delete(um.onlineUsers, UserID)
}

// GetAllOnLineUsers 返回所有
func (um *UserMgr) GetAllOnLineUsers() map[int]*UserProcess {
	return um.onlineUsers
}

// GetOnLineUserByID 返回指定用户
func (um *UserMgr) GetOnLineUserByID(userID int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userID]
	if !ok {
		err = fmt.Errorf("用户%d不在线", userID)
		return
	}
	return
}
