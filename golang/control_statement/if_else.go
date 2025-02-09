package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	// if 语句跟其他编程语言相同
	weekday := time.Now().Weekday()
	if weekday == 5 {
		fmt.Println("今天周五")
	} else if (weekday == 0) {
		fmt.Println("周末好快")
	} else {
		fmt.Println("打工啦")
	}
	var i int = 1
	/*if i {
	if 条件只能是一个 bool 值或者返回 bool 值得表达式，而不能是 int
	}*/
	if i > 0 {
		fmt.Println("大于0")
	}
	/**
	 * if 可以包含一个初始化语句（如：给一个变量赋值）。这种写法具有固定的格式（在初始化语句后方必须加上分号）
	 * 使用简短方式 := 声明的变量的作用域只存在于 if 结构中（在 if 结构的大括号之间，如果使用 if-else 结构则在 else 代码块中变量也会存在）。
	 * 如果变量在 if 结构之前就已经存在，那么在 if 结构中，该变量原来的值会被隐藏。最简单的解决方案就是不要在初始化语句中声明变量*
	 */
	var val int = 5
	if val := 10; val > 5 {
		fmt.Println("大于5")
	}
	fmt.Printf("val 的值: %d\n", val)

	// Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败
	// 当不使用 true 或 false 的时候，也可以使用一个 error 类型的变量来代替作为第二个返回值：成功执行的话，error 的值为 nil，否则就会包含相应的错误信息（Go 语言中的错误类型为 error: var err error）。
	// 这样一来，就很明显需要用一个 if 语句来测试执行结果；
	// 由于其符号的原因，这样的形式又称之为 comma,ok 模式（pattern）。
	// if 语句 与 多返回值函数的错误关系
	f, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
	/**
	 * 或者可以这样使用
	 * if value, ok := readData(); ok {
	   …
	 * }
	 */
}


