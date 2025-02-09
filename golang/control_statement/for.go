package main

import "fmt"

func main() {
	// Go 只有 for 语句, 没有while {} do {} while 语句

	// for 普通形式
	for i:=0; i < 5; i++ {
		fmt.Printf("this is the %d iteration\n", i)
	}

	// for 基于条件判断的迭代形式
	var i int = 5
	for i >= 0 {
		i = i - 1
		fmt.Printf("the variable i is now: %d\n", i)
	}

	// for 无限循环 形式 for ;; {} 或者 for true {}  或者 for {}
	var j int = 1
	for {
		if j >= 2 {
			break // 跳出循环
		}
		fmt.Printf("the variable j is now: %d\n", j)
		j++
	}

	// for-range 结构
	// 它可以迭代任何一个集合。语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。一般形式为：for ix, val := range coll { }。
	// val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值, 如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值。
	// 一个字符串是 Unicode 编码的字符（或称之为 rune）集合，因此您也可以用它迭代字符串
	str := "Go is a beautiful language"
	for pos, char := range str {
		fmt.Printf("character on position %d is : %c \n", pos, char)
	}
	str1 := "Go"
	// for-range 也可以 for ix := range coll {} 或者 可以 _ 把 ix 不要, 迭代 val
	for _, char := range str1 {
		fmt.Printf("character: %c\n", char)
	}
	// 一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。
	// 但在 switch 或 select 语句中，break 语句的作用结果是跳过整个代码块，执行后续的代码。

	// 关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件
}
