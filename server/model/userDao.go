package model

import (
	"encoding/json"
	"fmt"
	"sms/common/message"

	"github.com/garyburd/redigo/redis"
)

var (
	MyUserDao *UserDao
)

// UserDao 完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// NewUserDao 工厂模式，创建UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//getUserByID 根据用户ID，返回User实例
func (ud *UserDao) getUserByID(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNotExists
		}
		return
	}
	user = &message.User{}
	// 将user反序列化成user对象
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err")
	}
	return
}

// Login 登陆校验
func (ud *UserDao) Login(userID int, userPwd string) (user *message.User, err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserByID(conn, userID)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ErrorUserPwd
		return
	}
	fmt.Printf("用户%s成功登陆\n", user.UserName)
	return
}

// Register 注册校验
func (ud *UserDao) Register(user *message.User) (err error) {
	conn := ud.pool.Get()
	defer conn.Close()
	_, err = ud.getUserByID(conn, user.UserID)
	if err == nil {
		err = ErrorUserExists
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	_, err = conn.Do("HSet", "users", user.UserID, string(data))
	if err != nil {
		fmt.Println("保存注册用户失败 ", err)
	}
	return
}
