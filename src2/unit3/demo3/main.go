package main

import "fmt"

func main() {
	/*
		//求1-100的和，当第一次超过300时停止输出
		var sum int = 0
		//标签
	lable:
		for i := 0; i <= 100; i++ {
			sum += i
			fmt.Println(sum)
			if sum >= 300 {
				break lable
			}
		}
	*/
	//1-100中被6整除的数字
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
