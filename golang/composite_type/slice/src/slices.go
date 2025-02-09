package main

import "fmt"

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr1 = [3]int{0, 1, 2}
	var slice1 []int = arr1[1:3]
	slice1[0] = 3
	fmt.Println(arr1)
	fmt.Println(slice1)

	var arr2 = [5]int{0, 1, 2, 3, 4}
	fmt.Printf("arr2 之和为: %d\n", sum(arr2[:])) // 10

	var slice2 []int = make([]int, 10)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = 5 * i
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("Slice2 at %d is %d\n", i, slice2[i])
	}

	// 切片在达到容量上限后可以扩容。改变切片长度的过程称之为切片重组
	// 当我们在一个切片基础上重新划分一个切片时，新的切片会继续引用原有切片的数组
	slice3 := make([]int, 1, 2)
	slice3[0] = 1
	fmt.Printf("长度: %d\n", len(slice3))
	slice3 = slice3[0:len(slice3)+1]
	fmt.Printf("长度: %d\n", len(slice3))
	fmt.Println(slice3)

	// 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom)
	fmt.Printf("copied %d elements\n", n)

	// func append(s[]T, x ...T) []T 其中 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片；
	// 追加的元素必须和原切片的元素同类型。
	// 如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。
	// 因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}