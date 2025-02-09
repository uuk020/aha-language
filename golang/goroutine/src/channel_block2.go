package main

import (
	"fmt"
	"time"
)

func main()  {
	stream := pump1()
	go suck1(stream)
	time.Sleep(1e9)
}

func pump1() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck1(ch chan int)  {
	for {
		fmt.Println(<-ch)
	}
}