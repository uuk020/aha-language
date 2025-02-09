package main

import "fmt"

func main() {
	var a int
	var b int32
	a = 15
	// Go 中不允许不同类型之间的混合使用，但是对于常量的类型限制非常少，因此允许常量之间的混合使用
	// b = a + a
	b = b + 5
	fmt.Printf("a 的类型: %T\n", a)
	fmt.Printf("b 的类型: %T\n", b)
}
