package main

import "fmt"

func main() {
	arr2 := [2]int{0, 1}
	arr3 := arr2
	arr3[0] = 1
	fmt.Printf("arr2 的index: 0 位置: %d\n", arr2[0]) // 0
	fmt.Printf("arr3 的index: 0 位置: %d\n", arr3[0]) // 1

	arr4 := [1]int{1}
	arr5 := &arr4
	arr5[0] = 1
	fmt.Printf("arr4 的index: 0 位置: %d\n", arr4[0]) // 1
	fmt.Printf("arr5 的index: 0 位置: %d\n", arr5[0]) // 1

	var arr6 = new([2]int)
	arr6[0] = 1
	arr7 := arr6
	arr7[1] = 3
	fmt.Printf("arr6 的index: 1 位置: %d\n", arr6[1]) // 3
	fmt.Printf("arr7 的index: 1 位置: %d\n", arr7[1]) // 3

	var arr8 = new([2]int)
	arr9 := *arr8
	arr9[1] = 4
	fmt.Printf("arr6 的index: 1 位置: %d\n", arr8[1]) // 0
	fmt.Printf("arr7 的index: 1 位置: %d\n", arr9[1]) // 4

	var arr1 [5]int

	for i:=0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i:=0; i < len(arr1); i++ {
		fmt.Printf("array at index %d is %d\n", i, arr1[i])
	}

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}


}
