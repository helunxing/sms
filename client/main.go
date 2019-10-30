package main

import (
	"fmt"
	"net"
	"os"
)

var conn net.Conn

func init() {
	var err error
	conn, err = net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
}

func main() {
	defer conn.Close()
	// 接受选项，判断是否继续显示菜单
	var key int
	var loop = true
	for loop {
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
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	// 根据用户的输入，显示新的提示信息
	if key == 1 {
		for {
			var userID int
			var userPwd string
			fmt.Println("请输入用户id：")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			err := login(userID, userPwd)
			if err != nil {
				fmt.Printf("登陆失败，原因：%s\n", err.Error())
				fmt.Println("继续？y或Y继续，2注册，其他则退出：")
				var chr string
				fmt.Scanf("%d\n", &chr)
				if !(chr == "y" || chr == "Y") {
					break
				}
				if chr == "2" {
					fmt.Println("注册逻辑")
				}
			} else {
				fmt.Println("登陆成功")
				break
			}
		}
	} else if key == 2 {
		fmt.Println("注册逻辑")
	}
}
