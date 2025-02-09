
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	stack := make([]*tree.Tree, 0)
	for len(stack) > 0 || t != nil {
		for t != nil {
			stack = append(stack, t)
			t = t.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ch <- node.Value
		t = node.Right
	}
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 && ok2 {
			//两个信道都在开启状态
			if v1 != v2 {
				return false
			}
		} else if ok1 || ok2 {
			//一个信道已关闭，说明两棵树节点数不同
			return false
		} else {
			//两个信道同时关闭，说明两棵树节点数相同
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
