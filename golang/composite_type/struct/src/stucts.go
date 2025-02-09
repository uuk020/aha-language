package main

import (
	"fmt"
	"strings"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person)  {
	p.firstName = strings.ToUpper(p.firstName) // 像 p.lastName=strings.ToUpper(p.firstName) 这样给结构体字段赋值，没有像 C++ 中那样需要使用 -> 操作符，Go 会自动做这样的转换
	p.lastName = strings.ToUpper(p.lastName) // 也可以用 (*p).lastName = strings.ToUpper((*p).lastName)
}

func main()  {
	ms := new(struct1)
	fmt.Println(ms) // &{0, 0, } 的指针
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Wythe"
	fmt.Println(ms) // &{10, 15.5, Wythe}

	ms1 := struct1{i1: 1, f1: 10.4, str: "nathan"}
	fmt.Println(ms1) // {1, 10.4, Nathan}

	var pers1 Person
	pers1.firstName = "Huang"
	pers1.lastName = "Wiie"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	pers2 := &Person{"Huang", "Wiie"}
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	pers3 := new(Person)
	(*pers3).firstName = "Huang"
	(*pers3).lastName = "Wiie" // 也可以解引用方式来设置值
	fmt.Printf("The name of the person is %s %s\n", (*pers3).firstName, (*pers3).lastName)
}
