package model

import (
	"encoding/json"
	"fmt"

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
func (ud *UserDao) getUserByID(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNotExists
		}
		return
	}
	user = &User{}
	// 将user反序列化成user对象
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err")
	}
	return
}

// Login 登陆校验
func (ud *UserDao) Login(userID int, userPwd string) (user *User, err error) {
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
