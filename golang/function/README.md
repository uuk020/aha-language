## 函数

函数可以说是开发大型软件的基石，也是封装代码的基本单位。在一些比较老的不支持 oop 的编程语言中，正是一个个函数构建起来大型软件

### 定义函数

go 定义一个函数比较简单，语法如下
```
// optionalParameters 是 (param1 type1, param2 type2 ...) 这种形式
func functionName(optionalParameters) optionalReturnType {
  body
}
```

go 函数参数与返回值
 
函数不支持默认参数, 可通过传递零值并且代码里判断是否是零值来实现，另一种是通过传递一个结构体来实现

Go 默认使用按值传递来传递参数，也就是传递参数的副本。
函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量

如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加 & 符号，比如 &variable）传递给函数，
这就是按引用传递，比如 Function(&arg1)，此时传递给函数的是一个指针。如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；我们可以通过这个指针的值来修改这个值所指向的地址上的值。

**指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递。**

```go

func sum(a int, b int) int {
	return a + b
}

// 多个参数一致, 可以只写一个类型说明
func sum1(a, b int) int {
	return a + b
}

// 可变参数 用三个省略号来实现
// 注意可变参数其实被包装成了一个 slice
func sum3(init int, vals ...int) int {
	sum := init
	//vals is []int
	for _, val := range vals {
		sum += val
	}
	return sum
}
```

函数能够接收参数供自己使用，也可以返回零个或多个值（我们通常把返回多个值称为返回一组值）

```go
func sum1(a, b int) {
	return a + b
}

// 多返回值用括号包起来
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名的返回值 
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}
```

**当一个函数在其函数体内调用自身，则称之为递归**
```go
func fib(n int) int {
    if n < 2 {
        return n
    }
    return fib(n-1) + fib(n-2)
}
```

### defer 延迟语句

关键字 defer 允许我们推迟到函数返回之前一刻才执行某个语句或函数

> Any functions deferred by F are executed before F returns to its caller.

```go
func test() {
	fmt.Printf("test\n")
}

// 打印 main test
func main() {
	defer test()
	fmt.Printf("main\n")
}
```

> A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking. 

defer 执行时间分别是:
- 包裹defer的函数返回时
- 包裹defer的函数执行到末尾时
- 所在的goroutine发生panic时

defer 执行顺序: 按照"栈"形式, 先进后出

```go
func test1() {
	fmt.Printf("test1 \n")
}

func test2() {
	fmt.Printf("test2\n")
}

// 打印 main test2 test1
func main() {
	defer test1()
	defer test2()
	fmt.Printf("main\n")
}
```

defer 函数调用与参数: defer 调用的函数，参数的值在 defer 定义时就确定了

```go
func printArg(a int) {
    fmt.Printf("打印 a : %d\n", a)
}


func main() {
    var a int = 6
    defer printArg(a) // 打印出为 6 
    a++
    fmt.Printf("main 打印出 a 为: %d\n", a) // 打印出为 7
}

```

defer 关键字与循环

```go
func printArg(a int) {
	fmt.Printf("打印 a : %d\n", a)
}

func main() {
	for i := 0; i < 3; i++ {
	    defer printArg(i) // 循环执行 defer 语句 3 次, 根据先进后出, 打印出 2, 1, 0	
    }
}
```

参考资料
- [深入理解 defer 关键字](https://juejin.cn/post/6886710490530054158)
- [理解 Go 语言 defer关键字的原理](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/)

### 内置函数

Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作，例如：len、cap 和 append，或必须用于系统级的操作，
例如：panic。因此，它们需要直接获得编译器的支持

|  名称   | 说明  |
|  ----  | ----  |
| close  | 用于管道通信 |
| len、cap  | len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map） |
| new、make  | new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new (type)、make (type)。new (T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。它也可以被用于基本类型：v := new(int)。make (T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作）new () 是一个函数，不要忘记它的括号 |
| copy、append | copy、append |
| panic、recover  | 两者均用于错误处理机制 |
| print、println  | 底层打印函数，在部署环境中建议使用 fmt 包 |
| complex、real、imag  | 用于创建和操作复数 |

### 函数作为参数

go 里边函数其实也是『一等公民』，函数本身也是一种类型，所以我们可以定义一个函数然后赋值给一个变量; 还可以作为其他函数进行传递, 然后在其它函数内调用执行. 

```go
func funPrint(s string)  {
	fmt.Printf(s)
}

func main() {
	myPrint := funPrint // 函数类型赋值变量 myPrint
        myPrint("Hello World") // 调用 打印出 Hello World
        callback(1, add) // 打印出 1 + 2 之和: 3
}

func add(a, b int) {
	fmt.Printf("%d 和 %d 之和: %d\n", a, b, a + b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}
```

### 闭包

当我们不希望给函数起名字的时候, 可以使用匿名函数, 如 func(x, y int) int { return x + y }

这样的一个函数不能够独立存在（编译器会返回错误：non-declaration statement outside function body），但可以被赋值于某个变量，
即保存函数的地址到变量中：fplus := func(x, y int) int { return x + y }，然后通过变量名对函数进行调用：fplus(3,4)

也可以直接调用对匿名函数进行调用: func(x, y int) int { return x + y }(3, 4)

```go
func () {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum 值 :%d\n", sum)
}()
```

匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。另一种表示方式为：一个闭包继承了函数所声明时的作用域。
这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁

```go
func main() {
	str := "Hello World"
	func() {
	    str = "Hello Wythe"	
    }()
    fmt.Println(str)
}
```
在匿名函数总并没有声明或者定义 str 变量,但是可以直接操作, 根据闭包定义闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。另一种表示方式为：一个闭包继承了函数所声明时的作用域。
这种状态（作用域内的变量）都被共享到闭包的环境中. 闭包中共享(引用)了 main 函数中的 str 变量, 闭包对str变量改变也影响到 main 函数str  

```go
func main() {
    for i := 0; i < 3; i++ {
        defer func() {
            fmt.Printf("打印出 i 的值: %d\n", i) // 打印 2, 2, 2
        }()
    }
}
```
defer 延迟闭包执行, 闭包引用了 i 变量, 当循环结束后, i 的值变成了 2 . main 函数执行完毕之前执行了 defer , 打印出 2

defer 的时候明确介绍了 defer 后变量是保留它在 defer 时的值，而不会被 defer 之后的代码所改变。但是在闭包这边这个看起来不太适用，其实是适用的，只是闭包是引用了这个变量，
也就是说，在 defer 时被保留下来的是这个变量的地址，后续代码改变的不是地址，而是这个地址存储的值，所以后续代码对这个变量的操作，都会反应到这个 defer 中

可以在闭包调用中传入变量 i , 通过值传递来解决
```go
func main() {
    for i := 0; i < 3; i++ {
        defer func(a int) {
            fmt.Printf("打印出 i 的值: %d\n", a) // 打印 2, 1, 0
        }(i)
    }
}
```

参考资料:
- [Go 语言中的闭包](http://www.imooc.com/wiki/golesson/goclosure.html)
- [Go 使用 defer 函数 要注意的几个点](https://juejin.cn/post/6844904025708576776)

### 方法

Go 方法是作用在接收者 (receiver) 上的一个函数, 接收者是某种类型的变量, 因此方法是一种特殊类型的函数

接收者类型可以是 (几乎) 任何类型, 不仅仅是结构体类型: 任何类型都可以有方法, 甚至可以是函数类型, 可以是int, string, bool或者数组的别名类型. 
但是接收者不能是一个接口类型, 因为接口是一个抽象定义, 但是方法却是具体实现.

最后接收者不能是一个指针类型, 但是它可以是任何其他允许类型的指针

一个类型加上它的方法等价于面向对象中的一个类。一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，
唯一的要求是：它们必须是同一个包的。

因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法。但是如果基于接收者类型，是有重载的：
具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的

```
func (a *denseMatrix) Add(b Matrix) Matrix
func (a *sparseMatrix) Add(b Matrix) Matrix 
```

别名类型不能有它原始类型上已经定义过的方法

```
func (recv receiver_type) methodName(parameter_list) (return_value_list) { .... }
```

在方法名之前，func 关键字之后的括号中指定 receiver。
如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 object.name 选择器符号：recv.Method1()。
如果 recv 一个指针，Go 会自动解引用。
如果方法不需要使用 recv 的值，可以用 _ 替换它

函数与方法区别

- 函数将变量作为参数：Function1(recv)
- 方法在变量上被调用：recv.Method1()

在接收者是指针时，方法可以改变接收者的值（或状态），这点函数也可以做到（当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态）


鉴于性能的原因，recv 最常见的是一个指向 receiver_type 的指针（因为我们不想要一个实例的拷贝，如果按值调用的话就会是这样），
特别是在 receiver 类型是结构体时，就更是如此了。 如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。否则，就在普通的值类型上定义方法。 

当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。
这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果

有两种方法来实现在类型中嵌入功能：
A：聚合（或组合）：包含一个所需功能类型的具名字段。
B：内嵌：内嵌（匿名地）所需功能类型

```go
package main

import (
    "fmt"
)

type Log struct {
    msg string
}

type Customer struct {
    Name string
    log  *Log // Log 指针类型
}

func main() {
    c := new(Customer)
    c.Name = "Barak Obama"
    c.log = new(Log)
    c.log.msg = "1 - Yes we can!"
    // shorter
    //c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
    // fmt.Println(c) &{Barak Obama 1 - Yes we can!}
    c.Log().Add("2 - After me the world will be a better place!")
    //fmt.Println(c.log)
    fmt.Println(c.Log())

}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}

func (l *Log) String() string {
    return l.msg
}

func (c *Customer) Log() *Log {
    return c.log
}
```