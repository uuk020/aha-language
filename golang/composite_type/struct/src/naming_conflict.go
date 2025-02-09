package main

import "fmt"

type T struct {a int}

type R struct {a, b int}

type C struct {T; R}

type D struct {R; b float32}

func main() {
	var c C
	c = C{T{1}, R{2, 3}}
	// 没有办法来解决这种问题引起的二义性，必须由程序员自己修正
	// 具体名称
	fmt.Println("c.a 是 c.A.a 还是 c.B.a: ", c.T.a) // c.a 会报错 ambiguous selector c.a
	var d D
	d = D{R{1, 2}, 3.5}
	// 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
	// 如果想要内层的 b 可以通过 d.R.b 得到
	fmt.Println("d.b打印出来的是: ", d.b) // 3.5
}
