package main

import (
	"fmt"
)

func test(num int) {
	fmt.Println(num)
}

// 定义一个函数，把另一个函数作为形参：
func test2(num1 int, num2 float32, testFunc func(int)) {
	_ = num1
	_ = num2
	_ = testFunc
	fmt.Println("-----test2")
}

type myFunc func(int)

func test3(num1 int, num2 float32, testFunc myFunc) {
	_ = num1
	_ = num2
	_ = testFunc
	fmt.Println("-----test3")
}

// 定义一个函数求和，差 （传统写法返回值和返回类型一一对应，顺序不能差)
func test4(num1 int, num2 int) (int, int) {
	result1 := num1 + num2
	result2 := num1 - num2
	return result1, result2
}

// 升级写法 对函数返回值命名，里面的顺序就无所谓了，顺序不用对应
func test5(num1 int, num2 int) (sum int, sub int) {
	sub = num1 - num2
	sum = num1 + num2
	return
}
func main() {
	//函数也是一种数据类型，可以赋值给一个变量
	a := test                                    //变量就是一个函数类型的变量
	fmt.Printf("a的类型：%T,test的类型:%T \n", a, test) //a的类型：func(int),test的类型:func(int)

	a(10) //等价于test(10)
	//调用test2函数
	test2(10, 3.14, test)
	test2(10, 3.14, a)
	fmt.Println()

	//自定义数据类型：给int类型取了个别名叫myInt类型
	type myInt int
	var num3 myInt = 30
	var num4 int = 20
	//虽然是别名，但是在go中编译的时候还是认为myInt和int不是一种数据类型，num4=num3时要把num3转为int
	num4 = int(num3)
	fmt.Println(num3, num4)
	test3(10, 4.1, a)

	c, d := test4(30, 15)
	e, f := test5(45, 15)
	fmt.Printf("test4---Sum: %v  Sub: %v \n", c, d)
	fmt.Printf("test5---Sum: %v  Sub: %v \n", e, f)

}
