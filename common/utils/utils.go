package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sms/common/message"
)

// WritePkg 向连接中写入数据包
func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送长度给对方
	fmt.Printf("给%s发送数据\n", conn.RemoteAddr().String())
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	var n int
	n, err = conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	// 再正式发送数据
	n, err = conn.Write(data)
	// var sentl uint32
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	return
}

// ReadPkg 从连接中接收数据包
func ReadPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	n, err := conn.Read(buf[:4])
	// 此处：A处
	if n != 4 || err != nil {
		return
	}
	// 此句话如果上移到A处，则最后断开时，会再输出一次
	fmt.Printf("自%s读取数据\n", conn.RemoteAddr().String())
	// fmt.Printf("字节数为%d\n", n)

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
