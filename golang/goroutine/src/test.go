package main

import (
	"fmt"
	"runtime"
)

func goroutine1() {
	defer fmt.Println("it is defer content")
	for i := 0; i < 5; i++ {
		runtime.Goexit() // Goexit可以中止当前协程的执行，但是不影响其他协程的正常执行,在当前协程中止之前，会执行所有的defer后的内容
		fmt.Println("goroutine1 is running....")
	}
}

func goroutine2(input int) string  {
	for i := 0; i < input; i++ {
		fmt.Println("goroutine2 is running, input is", input)
	}
	return "it is goroutine2"
}

func main()  {
	go goroutine1()
	go goroutine2(5)
	for i := 0; i < 10; i++ {
		runtime.Gosched() // Gosched会使当前协程暂时放弃cpu执行权，让其他协程先执行，之后会重新自动获得执行权
		fmt.Println("main is running....")
	}
}