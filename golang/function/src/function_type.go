// 函数类型
package main

import "fmt"

func funPrint(s string)  {
	fmt.Printf(s)
}

func main() {
	myPrint := funPrint
	myPrint("Hello World\n")
	callback(1, add)
}

func add(a, b int) {
	fmt.Printf("%d 和 %d 之和: %d\n", a, b, a + b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
