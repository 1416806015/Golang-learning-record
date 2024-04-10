package main

import (
	"fmt"
)

func main() {
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	var num3 float32 = 314e-2
	fmt.Println(num3)
	var num4 float32 = 314e+2
	fmt.Println(num4)
	num5 := 3.14222233333
	fmt.Printf("%T", num5)
}
