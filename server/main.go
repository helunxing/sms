package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器监听9999端口")
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 8096)
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Printf("读到的l=%d\n", buf[:4])
}
