package main

import (
	"fmt"
)

func main() {

	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)
	switch score / 10 {
	case 10:
		fmt.Println("你的等级为S")
	case 9:
		fmt.Println("你的等级为A")
	case 8:
		fmt.Println("你的等级为B")
	case 7:
		fmt.Println("你的等级为C")
	case 6:
		fmt.Println("你的等级为D")
	default:
		fmt.Println("你的成绩不合格")
	}
}
