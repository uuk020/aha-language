package main

import (
	"fmt"
	"math"
)

type TwoInts struct {
	a, b int
}

type IntVector []int

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

type NamedPoint struct {
	Point
	name string
}

// 覆写方法 和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法
func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}

func main()  {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("两数之和: %d\n", two1.AddThem())
	fmt.Println("结构体 TwoInts: ", two1)
	fmt.Printf("两数之和: %d\n", two1.ChangeAndAddThem())
	fmt.Println("结构体 TwoInts: ", two1)
	fmt.Println(IntVector{1, 2, 3}.Sum())

	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}

// AddThem 不用指针地址
func (tn TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) ChangeAndAddThem() int {
	tn.a = tn.a + 10
	return tn.a + tn.b
}

func (v IntVector) Sum() (s int)  {
	for _, x := range v {
		s += x
	}
	return
}
