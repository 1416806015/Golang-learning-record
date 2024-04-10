package main

import "fmt"

func exchangeNum(num1 *int, num2 *int) {
	//对地址对应的变量进行改变数值
	var t int = *num1
	*num1 = *num2
	*num2 = t
}

func main() {

	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前：num1 = %v num2 = %v \n", num1, num2)
	exchangeNum(&num1, &num2)
	fmt.Printf("交换后：num1 = %v num2 = %v \n", num1, num2)
}
