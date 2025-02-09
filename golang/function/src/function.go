package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sum1(a, b int) int {
	return a + b
}

func sum3(init int, vals ...int) int {
	sum := init
	//vals is []int
	for _, val := range vals {
		sum += val
	}
	return sum
}

func changeVariable(a int) int {
	a = a + 5
	return a
}

func changeVariable1(a *int) int {
	*a = *a + 5
	return *a
}

func concat(a int) int{
	if a == 0 {
		a = 5
	}
	return a
}

func sum2(a, b int) (res int) {
	res = a + b
	return
}


func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func main() {
	// 函数调用跟其他语言无差 函数名称加上括号(有形参数则传入实参)
	fmt.Printf("调用函数 sum : %d\n", sum(1, 2))
	fmt.Printf("调用函数 sum1: %d\n", sum1(2, 3))
	fmt.Printf("调用可变参数 sum3: %d\n", sum3(1, 2, 3))
	var a int = 1
	var b = changeVariable(a)
	fmt.Printf("值传递调用: %d\n", b)
	fmt.Printf("变量 a 的值: %d\n", a)
	var t int = 1
	var p = changeVariable1(&t)
	fmt.Printf("值传递调用: %d\n", p)
	fmt.Printf("变量 t 的值: %d\n", t)
	fmt.Printf("默认参数: %d\n", concat(0))
	fmt.Printf("返回值变量: %d\n", sum2(3, 4))
	var i, j = getX2AndX3(2)
	fmt.Printf("多个返回值: %d, %d\n", i, j)
}
