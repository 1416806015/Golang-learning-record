package main

import "fmt"

func main() {
	/*
		var sum int = 0
		for i := 0; i <= 5; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/
	var str string = "hello golang"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	/*     for range 结构
	for key,val := range coll{

	}
	*/
	//对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体数值被value接收
	//遍历对字符进行遍历
	for i, value := range str {
		fmt.Printf("索引为：%d,具体值为:%c \n", i, value)
	}

}
