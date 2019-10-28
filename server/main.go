package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"sms/common/message"
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

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		// err = errors.New("read pkg header error")
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	n, err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}
func process(conn net.Conn) (mes message.Message, err error) {
	defer conn.Close()
	for {
		mes, err = readPkg(conn)
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Println("客户端已断开连接")
			default:
				fmt.Println("readPkg fail ", err)
			}
			return
		}
		fmt.Println("mes=", mes)
	}
}
