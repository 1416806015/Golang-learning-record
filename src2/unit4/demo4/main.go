package main

import "fmt"

//函数功能：求和
//函数的名字：getSUm 参数为空
//getSUm函数返回值是一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

//闭包：返回的匿名函数+匿名函数以外的变量sum
//匿名函数引用的那个变量会一直存在内存中，可以一直使用(对内存消耗大)
func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println("---------------------")
	fmt.Println(getSUm01(0, 1))
	fmt.Println(getSUm01(1, 2))
	fmt.Println(getSUm01(2, 3))
}

//不使用闭包时：我想保留的值不可重复使用
func getSUm01(sum int, num int) int {
	sum = sum + num
	return sum
}
