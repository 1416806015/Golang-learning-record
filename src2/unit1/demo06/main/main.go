package main

import (
	"fmt"
)

func main() {
	//指针   1.  & 取地址 2.  * 取值
	var age int = 18
	fmt.Println("age 的地址为:", &age) //age的地址 0xc000096068
	//定义指针变量 var声明一个变量 ptr指针的名字 *int是ptr对应的指针类型（指向int类型的指针）&age为一个地址，即ptr的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr 本身的存储空间地址为：", &ptr)
	//*ptr为取地址为0xc000096068的值 18
	fmt.Println("ptr 指向的值为：", *ptr)
}
