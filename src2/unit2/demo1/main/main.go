package main

import (
	"fmt"
	"unit2/demo1/test"
)

func main() {

	var num int = 10
	fmt.Println(num)
	var ptr *int = &num
	// var ptr *int = num        ps:指针变量接收的一定是地址值
	// var ptr *float32 = &num   ps: *float32意味着这个指针指向float32类型的数据，但&num对应的是int类型的不可以

	*ptr = 20
	fmt.Println(num)

	//调用自己创建的包
	fmt.Println(test.StuNo)
}
