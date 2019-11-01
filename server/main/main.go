package main

import (
	"fmt"
	"net"
	"sms/server/model"
	"time"
)

func main() {
	// 两个初始化有顺序
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器监听9999端口")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	processor := Processor{
		Conn: conn,
	}
	fmt.Printf("新协程开始处理与%s的连接\n", conn.RemoteAddr().String())
	err := processor.process()
	if err != nil {
		fmt.Printf("处理与%s连接的协程异常退出：%s\n", conn.RemoteAddr(), err)
	} else {
		fmt.Printf("处理与%s连接的协程正常退出\n", conn.RemoteAddr())
	}
	return
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)

}
