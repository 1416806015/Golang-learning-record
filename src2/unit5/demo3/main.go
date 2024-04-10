package main

import "fmt"

//new 函数： 分配内存，主要用来分配值类型（int，float，bool，string，数组和结构体）

func main() {
	//new函数的实参是一个类型而不是具体数值，new函数返回值是对应类型的指针 num： *int
	num := new(int)
	fmt.Printf("num的类型：%T，num的值：%v，num的地址：%v，num指针指向的值：%v\n", num, num, &num, *num)

}
