package main

import (
	"fmt"
	"time"
)

func calculate(a, b int) int {
	return a + b
}

func main()  {
	// Go 语言中的 switch 结构使用上更加灵活。它接受任意形式的表达式
	// 类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式
	// 可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3
	// 一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
	// 因此，程序也不会自动地去执行下一个分支的代码。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。
	// 在 case ...: 语句之后，您不需要使用花括号将多行语句括起来，但您可以在分支中进行任意形式的编码。当代码块只有一行时，可以直接放置在 case 语句之后
	// 可选的 default 分支可以出现在任何顺序，但最好将它放在最后
	var num1 int = 100
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
	// switch 语句也跟 if 语句也可以包含一个初始化语句
	// switch 语句的第二种形式是不提供任何被判断的值（实际上默认为判断是否为 true），然后在每个 case 分支中进行测试不同的条件。当任一分支的测试结果为 true 时，该分支的代码会被执行
	switch weekday := time.Now().Weekday(); {
	case weekday == 0 && weekday == 6:
		fmt.Println("周末")
	case weekday >= 1 && weekday <= 5:
		fmt.Println("工作日")
	default:
		fmt.Println("上帝知道~")
	}

	result := calculate(1, 2)
	switch {
	case result < 0:
		fmt.Println("小于0")
	case result > 0 && result < 10:
		fmt.Println("在0和10之间")
	default:
		fmt.Println("大于10")
	}

	k := 6
	switch k {
	case 4: fmt.Println("was <= 4"); fallthrough
	case 5: fmt.Println("was <= 5"); fallthrough
	case 6: fmt.Println("was <= 6"); fallthrough
	case 7: fmt.Println("was <= 7"); fallthrough
	case 8: fmt.Println("was <= 8"); fallthrough
	default: fmt.Println("default case")
	}
}
