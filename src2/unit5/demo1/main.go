package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//统计字符串的长度按字节进行统计。
	str := "golang你好"     //在golang中，汉字是utf-8字符集，一个汉字占三个字节。
	fmt.Println(len(str)) //12

	//对字符串进行遍历
	//1、for-range
	for i, value := range str {
		fmt.Printf("索引为：%d 具体值为：%c\n", i, value)
	}
	//2、r :=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	//字符串转整数
	num, _ := strconv.Atoi("666")
	fmt.Println(num)
	//整数转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)
	//统计一下字符串里有几个子串
	count := strings.Count("golanggaskfga", "ga")
	fmt.Println(count)
	//不区分大小写的字符串比较
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)
	//区分大小写
	fmt.Println("hello" == "HELLO")
	//返回子串在字符串第一次出现的索引值，如果没有返回-1
	fmt.Println(strings.Index("golanggaskfga", "ga"))

}
