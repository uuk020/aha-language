package main

var a string

// GOG
func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O" // 局部变量声明和赋值
	print(a)
	f2()
}

func f2() {
	print(a)
}
