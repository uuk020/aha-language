package main

import (
	"fmt"
	"strings"
)

func main() {
	// HasPrefix 判断字符串 s 是否以 prefix 开头, 区分大小写
	str := "Hello World"
	fmt.Printf("%t\n", strings.HasPrefix(str, "H"))
	fmt.Printf("%t\n", strings.HasPrefix(str, "h"))

	// HasSuffix 判断字符串 s 是否以 suffix 结尾
	fmt.Printf("%t\n", strings.HasSuffix(str, "d"))
	fmt.Printf("%t\n", strings.HasSuffix(str, "D"))

	// 字符串包含关系, 区分大小写
	fmt.Printf("the strings str contain e : %t\n", strings.Contains(str, "e"))
	fmt.Printf("the strings str contain E : %t\n", strings.Contains(str, "E"))

	// 判断子字符串或字符在父字符串中出现的首次位置（索引）, -1表示字符串 是 不包含字符串 str
	fmt.Printf("index: %d\n", strings.Index(str, "e"))
	fmt.Printf("index: %d\n", strings.Index(str, "E"))
	fmt.Printf("index: %d\n", strings.Index(str, "l"))

	// 判断子字符串或支付在父字符串中出现的最后位置(索引), -1表示字符串 是 不包含字符串 str
	fmt.Printf("last index: %d\n", strings.LastIndex(str, "l"))

	//如果 ch 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位
	str2 := "您好, 世界"
	fmt.Printf("中文定位: %d\n", strings.IndexRune(str2, '世'))
	newStr := strings.Replace(str, "World", "Wythe", 6)
	fmt.Printf("new string: %s\n", newStr)
	fmt.Printf("number of H's in %s is: %d\n", str, strings.Count(str, "H"))
	var origS string = "Hi there!"
	var newS string
	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is %s\n", newS)

	//ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符
	var UpperStr = "GG"
	var lowerStr = strings.ToLower(UpperStr)

	// ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符
	fmt.Printf("The new string is %s\n", lowerStr)
	var UpperStr1 = strings.ToUpper(newStr)
	fmt.Printf("The new string is: %s\n", UpperStr1)
	str3 := " Hello World "
	/**
	 * strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号；
	 * 如果你想要剔除指定字符，则可以使用 strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉。该函数的第二个参数可以包含任何字符，
	 * 如果你只想剔除开头或者结尾的字符串，则可以使用 TrimLeft 或者 TrimRight 来实现
	 */
	trimSpaceStr := strings.TrimSpace(str3)
	fmt.Printf("have space string: %s\n", str3)
	fmt.Printf("no space string: %s\n", trimSpaceStr)

	str5 := "sZeros"
	trimStr := strings.Trim(str5, "s")
	fmt.Printf("trimed string: %s\n", trimStr)

	/**
	 * strings.Fields(s) 利用空白作为分隔符将字符串分割为若干块，并返回一个 slice 。如果字符串只包含空白符号，返回一个长度为 0 的 slice 。
	 * strings.Split(s, sep) 自定义分割符号对字符串分割，返回 slice 。
	 * 因为这 2 个函数都会返回 slice，所以习惯使用 for-range 循环来对其进行处理
	 * Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
	 */
	str6 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str6)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str7 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str7, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str8 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str8)

	// 函数 strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
	reader := strings.NewReader("abcde")
	fmt.Println(reader.Len())
	p := make([]byte, 2)
	// read() 从 [] byte 中读取内容
	c, _ := reader.Read(p)
	fmt.Print("读取的字节: ")
	fmt.Println(c)
	fmt.Print("返回未读的字符串长度: ")
	noReadLen := reader.Len()
	fmt.Println(noReadLen)
	fmt.Println(string(p))

	// ReadByte() 和 ReadRune() 从字符串中读取下一个 byte 或者 rune 会改变未读字符串长度。
	newReader := strings.NewReader("Hello")
	d, _ := newReader.ReadByte()
	fmt.Printf("未读的字符串长度: %d\n", newReader.Len())
	fmt.Printf("读取字节转换为字符: %c\n", d)
	otherReader := strings.NewReader("婷婷")
	t, size, _ := otherReader.ReadRune()
	fmt.Printf("未读的字符串长度: %d\n", otherReader.Len())
	fmt.Printf("读取字符编码转换为字符: %c , 编码长度: %v\n", t, size)
}
