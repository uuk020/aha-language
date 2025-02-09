package main

import "fmt"

func test()  {
	fmt.Printf("test\n")
}

func test1()  {
	fmt.Printf("test1\n")
}

func test2()  {
	fmt.Printf("test2\n")
}

func deferFunc() {
	fmt.Printf("defer func\n")
}

func returnFunc() int {
	fmt.Printf("return func\n")
	return 1
}

func printArg(a int) {
	fmt.Printf("打印 a : %d\n", a)
}

func main()  {
	//defer test()
	//defer test1()
	//defer test2()
	//defer deferFunc()
	//returnFunc()
	//fmt.Printf("main\n")
	//var a int = 6
	//defer printArg(a)
	//a++
	//fmt.Printf("main 打印出 a 为: %d\n", a)
	for i:=0; i < 3; i++ {
		defer printArg(i)
	}
}