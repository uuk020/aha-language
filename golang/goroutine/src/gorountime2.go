package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string)  {
	ch <- "深圳"
	ch <- "广州"
	ch <- "珠海"
}

func getData(ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s ", input)
	}
}