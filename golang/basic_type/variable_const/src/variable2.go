package main

import "fmt"

func main() {
	// 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
	// 像数组和结构这些复合类型也是值类型
	var i int = 1
	var j int
	//当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
	i = j
	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", j)
	//在 Go 语言中，指针属于引用类型，其它的引用类型还包括 slices，maps 和 channel。被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间
}
