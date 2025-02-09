package main

import (
	"fmt"
	"math"
)

type Square1 struct {
	side float32
}

type Circle struct {
	radius float32
}

type OtherShaper interface {
	OtherArea() float32
}

func main()  {
	var areaIntf OtherShaper

	//
	sq1 := new(Square1) //&Square1{5}
	sq1.side = 5

	areaIntf = sq1
	if t, ok := areaIntf.(*Square1); ok {
		fmt.Printf("areaIntf 类型是: %T\n", t)
	}

	var areaIntf1 OtherShaper
	ci1 := Circle{1}

	areaIntf1 = ci1

	if t, ok := areaIntf1.(Circle); ok {
		fmt.Printf("areaIntf1 类型是: %T\n", t)
	}
}

func (sq *Square1) OtherArea() float32 {
	return sq.side * sq.side
}

func (ci Circle) OtherArea() float32  {
	return ci.radius * ci.radius * math.Pi
}

