package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface {}

func main()  {
	var val Any
	val = 5
	fmt.Printf("val 的值为: %v\n", val)
	val = str
	fmt.Printf("val 的值为: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Wythe"
	pers1.age = 27
	val = pers1
	fmt.Printf("val 的值为: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("int 类型: %T\n", t)
	case string:
		fmt.Printf("string 类型: %T\n", t)
	case bool:
		fmt.Printf("bool 类型: %T\n", t)
	case *Person:
		fmt.Printf("Person的指针类型: %T\n", t)
	default:
		fmt.Printf("意外类型: %T\n", t)
	}
}
