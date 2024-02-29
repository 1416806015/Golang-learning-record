package main

import (
	"fmt"
	"unsafe"
)

var (
	num int = 19
	// num2 int = 100
)

func bey() {
	fmt.Println("bay")
}
func hey() {
	fmt.Println("hey")
}

func main() {
	// var age int = 18
	// fmt.Println("age =", age)
	// name, sex, age := "chen", "男", 18
	// fmt.Println(name, sex, age)
	bey()
	hey()
	fmt.Printf("num的类型是：%T", num)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num))

}
