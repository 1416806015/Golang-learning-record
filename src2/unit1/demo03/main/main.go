package main

import (
	"fmt"
)

func main() {
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '6'
	fmt.Println(c2)
	//汉字字符对应Unicode码值
	var c3 int = '陈'
	fmt.Println(c3)
	var c4 int = '陈'
	fmt.Printf("%c", c4)
	fmt.Println()
	//类型转换
	var n1 int = 100
	fmt.Printf("%T\n", n1)
	fmt.Println(n1)
	var n2 float32 = float32(n1)
	fmt.Println(n2)
	fmt.Printf("%T", n2)

	var n3 int64 = 12
	var n4 int8 = int8(n3) + 127 //编译通过，但是结果溢出
	//var n5 int8 = int8(n3) + 128    128溢出int8 表示范围 编译不通过

	fmt.Println("\n", n4)

}
