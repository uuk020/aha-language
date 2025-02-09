## 数组

数组是具有相同 **唯一类型** 的一组以编号且长度固定的数据项序列; 

这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。数组长度必须是一个常量表达式，并且必须是一个非负整数。
数组长度也是数组类型的一部分， 所以 [5] int 和 [10] int 是属于不同类型的

数组元素也是索引来读取或者修改, 从 0 开始. 元素的数目，也称为长度或者数组大小必须是固定的并且在声明该数组时就给出（编译时需要知道数组长度以便分配内存）；数组长度最大为 2Gb

声明格式:
```
var identifer [len]type
```

可以声明与初始化一起

```
var identifer = [len]type{value1, value2, ..., valueN}
identifer := [len]type{value1, value2, ..., valueN}
```

也可以省略长度, 用三个省略号表示, Go 自动计算长度

```
var identifer = [len]type{value1, value2, ..., valueN}
identifer := [...]type{value1, value2, ..., valueN}
```

也可部分赋值
```
var identifer = [len]type{0: value1}
```

多维数组

```
var identifer = [N][M]Type
```


例子:
```go
var arr1 [5]int
var arr2 = [...]int{1, 2}
```

arr1 每个元素都是一个整型值, 当声明数组时所有的元素都会被自动初始化为默认值0


由于索引的存在，遍历数组的方法自然就是使用 for 结构

- 通过 for 初始化数组项
- 通过 for 打印数组元素
- 通过 for 依次处理元素

Go 语言中数组是一种**值类型**, 把一组数组赋值给另外一个变量, 另一个变量修改数组中元素也不会对原数组有所影响

```go
arr2 := [2]int{0, 1}
arr3 := arr2
arr3[0] = 1
fmt.Printf("arr2 的index: 0 位置: %d\n", arr2[0]) // 打印 0
fmt.Printf("arr3 的index: 0 位置: %d\n", arr3[0]) // 打印 1
```

可以通过 new 函数创建, 而 new 返回的是s是其地址 *[5]int, 因此赋值给一个变量修改数组元素也会影响原数组的

```go
var arr6 = new([2]int)
arr7 := arr6
arr7[1] = 3
fmt.Printf("arr6 的index: 1 位置: %d\n", arr6[1]) // 3
fmt.Printf("arr7 的index: 1 位置: %d\n", arr7[1]) // 3
```

这样的结果就是当把一个数组赋值给另一个时，需要做解引用操作

```go
var arr8 = new([2]int)
arr9 := *arr8
arr9[1] = 4
fmt.Printf("arr8 的index: 1 位置: %d\n", arr8[1]) // 0
fmt.Printf("arr9 的index: 1 位置: %d\n", arr9[1]) // 4
```

将数组传递给函数: 由于数组是值类型, 因此传递函数是会拷贝数组, 所以把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
- 传递数组的指针使用数组的切片
- 使用数组的切片