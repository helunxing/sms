package main

import (
	"fmt"
)

// 登陆校验
func login(userID int, userPwd string) (err error) {
	fmt.Printf("id=%d pwd=%s\n", userID, userPwd)
	return nil
}
