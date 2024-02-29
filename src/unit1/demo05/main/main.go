package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string --> bool
	var s1 string = "true"
	var b bool
	//ParseBool这个函数的返回值有俩个：（value bool，err error）
	//value就是我们得到的布尔类型的数据，err出现的错误
	//我们只关注关注布尔类型得到的数据，而错误可以用 _ 直接忽略
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("s1对应的类型是 %T , s1 = %v \n", b, b)
	//string -->int64
	var s2 string = "19"
	var num int64
	num, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("num对应的类型是 %T , s1 = %v \n", num, num)
	//string -->float64
	var s3 string = "3.1415"
	var f1 float64
	f1, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("f1对应的类型是 %T , f1 = %v \n", f1, f1)
}
