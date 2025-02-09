package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("在 main()")
	go longWait()
	go shortWait()
	fmt.Println("睡眠 main")
	time.Sleep(10 * 1e9)
	fmt.Println("结束 main")
}

func longWait()  {
	fmt.Println("longWait 开始")
	time.Sleep(5 * 1e9)
	fmt.Println("longWait 结束")
}

func shortWait()  {
	fmt.Println("shortWait 开始")
	time.Sleep(2 * 1e9)
	fmt.Println("shortWait 结束")
}
