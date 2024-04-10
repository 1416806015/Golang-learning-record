package main

import "fmt"

func test1(arr *[3]int) {
	(*arr)[0] = 5
}

func main() {
	var arr1 = [3]int{7, 8, 9}
	test1(&arr1)
	fmt.Println(arr1)
}
