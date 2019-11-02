package main

import (
	"fmt"
	"os"
	"sms/client/process"
)

func main() {

	var userID int
	var userPwd string
	var userName string

	// 接收选项，判断是否继续显示菜单
	var key int
	for true {
		fmt.Println("——————欢迎登陆多人聊天系统——————")
		fmt.Println("\t\t 1 登陆聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择(1-3):")
		// 若不加换行将全部都被接收
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Printf("请输入用户id：")
			fmt.Scanf("%d\n", &userID)
			fmt.Printf("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			up := process.UesrProcess{}
			up.Login(userID, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Printf("请输入用户id：")
			fmt.Scanf("%d\n", &userID)
			fmt.Printf("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Printf("请输入昵称：")
			fmt.Scanf("%s\n", &userName)
			up := process.UesrProcess{}
			up.Register(userID, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
