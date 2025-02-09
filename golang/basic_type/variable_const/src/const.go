package main

import "fmt";

func main() {
	// 常量使用关键字 const 定义， 用于存储不会改变的数据，只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
	// 定义形式 const identifier [type] = value

	const Pi_type float64 = 3.14159
	fmt.Printf("%T", Pi_type) // float64
	fmt.Print("\n")

	// 可以省略类型说明符，因为编译器可以根据常量的值来推断其类型
	const Pi = 3.14159
	fmt.Printf("%T", Pi) // float64
	fmt.Print("\n")

	// 常量的值必须是能够在编译时就能够确定的；你可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。
	// 另外，常量也允许使用并行赋值的形式
	const Monday, Tuesday, Wednesday, Thursday, Friday = 1, 2, 3, 4, 5

	// 常量还可以用作枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)
	// 常量枚举，若下一行无赋值，则继承上一行的值
	const (
		Saturday = 6
		OtherSaturday
	)
	fmt.Println(OtherSaturday) // 则打印出来是6

	// iota 可以被用作枚举值，
	// 一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1
	// 在每遇到一个新的常量块或单个常量声明时， iota 都会重置为 0（ 简单地讲，每遇到一次 const 关键字，iota 就重置为 0 ）
	const (
		a = iota
		b
		c
	)
	fmt.Println(c) // 2
}
