package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str is %d\n", len(str))
	fmt.Printf("The number of characters in string str is %d\n", utf8.RuneCountInString(str))
	str2 := "你好，世界a"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of character in string str2 is %d\n", utf8.RuneCountInString(str2))
}
