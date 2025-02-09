package main

import "fmt"

func main()  {
	func () {
		sum := 0
		for i := 1; i < 10; i++ {
			sum += i
		}
		fmt.Printf("sum 值: %d\n", sum)
	}() // 打印 sum 值: 45

	str := "Hello World"
	func () {
		str = "Hello Wythe"
	}()
	fmt.Println(str)

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Printf("打印出 i 的值: %d\n", i)
		}()
	}
}
