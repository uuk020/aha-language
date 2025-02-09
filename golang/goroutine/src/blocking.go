package main
// 参考 https://segmentfault.com/q/1010000016868653
import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main()  {
	out := make(chan int)
	out <- 2
	go f1(out)
}
