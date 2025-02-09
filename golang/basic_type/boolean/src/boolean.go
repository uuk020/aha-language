package main

import "fmt"

func main() {
	type TZ int
	var aVar1 int = 10
	var aVar2 TZ = 10
	fmt.Printf("aVar1 与 aVar2 是否相等: %t\n", aVar1 == int(aVar2))
}
