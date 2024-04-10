package main

import "fmt"

func main() {
	/*var scores [5]int
	scores[0] = 98
	scores[1] = 88
	scores[2] = 45
	scores[3] = 65
	scores[4] = 55*/
	// 指定长度
	//var scores = [5]int{98, 87, 45, 65, 25}
	// 不指定长度
	//var scores = []int{98, 87, 45, 65, 25}
	//指定长度，通过索引值来初始化,1:1就表示，下标为1的元素，初始值为1. 下标默认初始值为0
	//var scores = []int{0: 98, 1: 87, 3: 45, 4: 65, 5: 25}
	//不指定长度，通过索引值进行初始化
	//var scores = []int{1: 98, 3: 99, 4: 23}

	var scores [5]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入第%d个学生成绩", i+1)
		fmt.Scanln(&scores[i])
	}
	for key, value := range scores {
		fmt.Printf("第%d个学生的成绩为%d\n", key+1, value)
	}
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg := sum / 5
	fmt.Printf("总成绩为：%v，平均成绩为：%d\n", sum, avg)

	fmt.Println("--------------------------------------")

}
