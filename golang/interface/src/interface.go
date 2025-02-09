package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

// 隐式地实现接口
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32  {
	return r.length * r.width
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	fmt.Println(areaIntf) // <nil>
	areaIntf = sq1 // 实现了这个接口的类型, 用这个类型定义的变量就可以给这个接口声明的变量赋值
	// 更短, 没有单独声明:
	// areaIntf := Shaper(sq1)
	// 或者:
	// areaIntf := sq1
	fmt.Printf("正方形面积为: %f\n", areaIntf.Area()) // 25.000000

	// 模拟多态
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	for n, _ := range shapes {
		fmt.Println("详情: ", shapes[n])
		fmt.Println("面积: ", shapes[n].Area())
	}
}