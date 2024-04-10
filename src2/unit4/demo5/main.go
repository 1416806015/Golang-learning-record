package main

import "fmt"

//defer使用场景：比如你想关闭某个使用的资源，在使用的时候直接随手defer，因为defer有延迟机制（函数执行完毕再执行的defer压入栈的语句）。
func main() {
	fmt.Println(add(30, 60))
}

func add(num1 int, num2 int) int {
	//在Golang中，程序遇到defer关键字，不会立马执行defer后的语句，而是将defer后的语句压入栈中，然后继续执行后面的语句
	defer fmt.Println("num1 = ", num1)
	defer fmt.Println("num2 = ", num2)
	num1 += 90
	num2 += 50
	var sum int = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
