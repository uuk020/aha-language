package main

import "fmt"

func main()  {
	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	fmt.Printf("map literal at \"one\" is: %d\n", mapList["one"])

	mapCreated := make(map[string]float32)
	mapCreated["key1"] = 3.14159
	fmt.Printf("map created at \"key1\" is %f\n", mapCreated["key1"])

	var mapAssigned map[string]int
	// mapAssigned 也是 mapList 的引用，对 mapAssigned 的修改也会影响到 mapLit 的值。
	mapAssigned = mapList
	mapAssigned["two"] = 3
	mapAssigned["three"] = 4
	fmt.Println(mapList)
	fmt.Println(mapAssigned)

	// 函数作为值的 map
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(mf) // 打印出来的是函数地址

	// 切片作为值的 map
	mp1 := make(map[int][]int)
	mp1[1] = []int{1, 2, 3}
	fmt.Println(mp1)
}
