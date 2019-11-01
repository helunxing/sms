package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sms/common/message"
)

// Transfer 结构体关联读写包方法
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

// ReadPkg 从连接中接收数据包
func (tf *Transfer) ReadPkg() (mes message.Message, err error) {

	n, err := tf.Conn.Read(tf.Buf[:4])
	// 此处：A处
	if n == 0 {
		err = errors.New("连接已断开")
	}
	if n != 4 || err != nil {
		return
	}
	// 此句话如果上移到A处，则最后断开时，会再输出一次
	fmt.Printf("自%s读取数据\n", tf.Conn.RemoteAddr().String())
	// fmt.Printf("字节数为%d\n", n)

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(tf.Buf[0:4])
	n, err = tf.Conn.Read(tf.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(tf.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}

	return
}

// WritePkg 向连接中写入数据包
func (tf *Transfer) WritePkg(data []byte) (err error) {
	// 先发送长度给对方
	fmt.Printf("给%s发送数据\n", tf.Conn.RemoteAddr().String())
	fmt.Printf("内容为%s", string(data))
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(tf.Buf[0:4], pkgLen)
	var n int
	n, err = tf.Conn.Write(tf.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	// 再正式发送数据
	n, err = tf.Conn.Write(data)
	// var sentl uint32
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	return
}
