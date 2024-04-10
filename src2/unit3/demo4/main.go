package main

import "fmt"

/*
func 函数名（形参列表）（返回值类型）{
	执行语句
	return 返回值列表
}

*/

func cal(num1 int, num2 int) int { //如果返回值类型就一个值的话，那么（）是可以省略不写的
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

func cal2(num1 int, num2 int) (int, int) { //如果返回值类型就一个值的话，那么（）是可以省略不写的
	var sum int = 0
	sum += num1
	sum += num2
	var result int = num1 - num2
	return sum, result
}

func main() {
	sum := cal(10, 20)
	sum1, result := cal2(10, 20)
	fmt.Println(sum, sum1, result)

}
