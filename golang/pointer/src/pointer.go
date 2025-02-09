package main

import "fmt"

func main()  {
	var i1 = 5
	fmt.Printf("i1 的值: %d, 内存地址为 %p\n", i1, &i1)
	var i = 5
	var intP *int
	intP = &i
	fmt.Printf("i的内存地址: %p\n", &i)
	fmt.Printf("intP的内存地址: %p\n", intP)
	*intP = 6
	fmt.Println(i)
	fmt.Println(*intP)
	fmt.Printf("改变后 i 的内存地址: %p\n", &i)
	fmt.Printf("改变后 intP的内存地址: %p\n", intP)

}
