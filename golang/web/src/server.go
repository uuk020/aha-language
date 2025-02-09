package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器启动...")
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("错误监听: ", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("错误接收: ", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("错误读取: ", err.Error())
			return
		}
		fmt.Printf("接收的数据: %v", string(buf[:len]))
	}
}
