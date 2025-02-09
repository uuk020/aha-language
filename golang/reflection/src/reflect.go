package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var x float64 = 3.4
	fmt.Println("类型: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("v 的值: ", v)
	fmt.Println("v 的类型: ", reflect.TypeOf(v))
	fmt.Println("kind 方法: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
