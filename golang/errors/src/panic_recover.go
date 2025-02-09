package main

import "fmt"

func badCall() {
	panic("中止")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("调用 badCall 之后")
}

func main()  {
	fmt.Printf("调用 test\r\n")
	test()
	fmt.Printf("调用完毕")
}
