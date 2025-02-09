package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	//go suck(ch1)
	//time.Sleep(1e9)
	fmt.Println(<-ch1)
}

func pump(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
