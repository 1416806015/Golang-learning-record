package main

import (
	"fmt"
	"strconv"
)

func main() {
	//方式一 fmt.Sprintf("%参数",表达式)
	var n1 int = 12
	var n2 float64 = 4.223
	var n3 bool = false
	var n4 byte = 'a'
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1对应的类型是 %T , s1 = %v \n", s1, s1)
	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("s2对应的类型是 %T , s1 = %v \n", s2, s2)
	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s3对应的类型是 %T , s1 = %v \n", s3, s3)
	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s3对应的类型是 %T , s1 = %v \n", s4, s4)

	//方式二 使用strconv包函数
	var n5 int = 18
	var s5 string = strconv.FormatInt(int64(n5), 10) //参数 第一个参数必须转为int64类型,第二个参数指定字面值的进制形式为十进制
	fmt.Printf("s5对应的类型是 %T , s5 = %q \n", s5, s5)

}
