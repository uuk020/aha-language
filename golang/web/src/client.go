package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("错误拨号: ", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("首先, 你的名字?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")
	for {
		fmt.Println("发送什么数据给服务端? 按 Q 退出")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
