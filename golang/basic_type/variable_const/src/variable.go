package main

import "fmt"

// 声明变量的一般形式是使用 var 关键字: var identifier type.
// 当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。记住，所有的内存在 Go 中都是经过初始化的。
var a int

// 也可以通过下面声明方式，这种因式分解关键字的写法一般用于声明全局变量
var (
	str string
	f   float64
)

// 一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为作用域
// 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
// 一般情况下，局部变量的作用域可以通过代码块（用大括号括起来的部分）判断。
func tempVar() {
	var tmp int
	fmt.Printf("局部变量: %d\n", tmp)
}

// 尽管变量的标识符必须是唯一的，但你可以在某个代码块的内层代码块中使用相同名称的变量，
// 则此时外部的同名变量将会暂时隐藏（结束内部代码块的执行后隐藏的外部同名变量又会出现，而内部同名变量则被释放），
// 你任何的操作都只会影响内部代码块的局部变量。
func tempVar1() {
	var b int8 = 1 // 局部变量 b
	if true {
		var b int8 = 2 // if 作用域里局部变量 b , 隐藏了函数 tempVar1 的局部变量 b
		fmt.Printf("if代码块的局部变量b: %d\n", b)
	}
	fmt.Printf("函数里的局部变量b: %d\n", b)
}

// 当你在函数体内声明局部变量时，应使用简短声明语法 :=
func tempVar2() {
	t := 3
	fmt.Printf("简短声明语法: %d\n", t)
}

func main() {
	fmt.Println(a)
	fmt.Println(str)
	str = "Hello World"
	fmt.Printf("字符串: %s\n", str[0:2])
	fmt.Printf("rune: %d\n", []rune(str)[0])
	fmt.Println(f)
	tempVar()
	tempVar1()
	tempVar2()
	// 声明变量并赋值
	var b bool = true
	fmt.Printf("%t\n", b)
}
