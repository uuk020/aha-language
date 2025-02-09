package main

import (
	"fmt"
	"log"
	"net/http"
)

type demoHandler struct {
	a string
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("在 HelloServer 处理")
	fmt.Fprintf(w, "Hello, " + req.URL.Path[1:])
}

func (obj *demoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("在 HelloServer 处理")
	fmt.Fprintf(w, "Hello, " + req.URL.Path[1:])
}

func main()  {
	//http.HandleFunc("/", HelloServer)
	//http.Handle("/", http.HandlerFunc(HelloServer))
	test := new(demoHandler)
	test.a = "Hello"
	http.Handle("/", test)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}