package main

import (
	"fmt"
	"time"
)

func main()  {
	var c = make(chan int)
	go source(c)
	go sink(c)
	time.Sleep(1e9)
}

func source(ch chan<- int)  {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func sink(ch <-chan int)  {
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}