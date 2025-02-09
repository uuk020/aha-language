package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("不能发现错误")

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return f * f, nil
}

func main()  {
	fmt.Printf("error: %v\n", errNotFound)
	res, err := Sqrt(-1.2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
