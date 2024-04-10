package main

import "fmt"

func main() {
	//键盘输入信息
	/*
			//法一：Scanln
			var name string
			fmt.Println("请输入名字:")
			//传入age地址的目的：在scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响
			fmt.Scanln(&name)

			var age int
			fmt.Println("请输入年龄:")
			fmt.Scanln(&age)

		fmt.Printf("名字：%v 年龄：%v", name, age)
	*/
	//法二：Scanf
	/*
		var name string
		var age int
		var score float32
		var isVIP bool
		fmt.Println("请输入信息，使用空格进行分隔:")
		fmt.Scanf("%s %d %f %t", &name, &age, &score, &isVIP)
		fmt.Printf("名字：%v 年龄：%v 成绩：%v VIP:%v", name, age, score, isVIP)

	*/

	var count int
	fmt.Println("请输入口罩个数：")
	fmt.Scanln(&count)
	if count < 30 {
		fmt.Println("对不起，口罩数量不足")
	} else {
		fmt.Println("口罩数量充足")
	}

}
