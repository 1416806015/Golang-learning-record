package main

import (
	"fmt"
	hengan "unit2/hen"
)

func main() {

	//make内置函数来创建切片
	//make的三个参数： 1.切片类型 2.切片长度 3.切片的容量
	slice := make([]int, 4, 20)
	fmt.Println(slice) //[0 0 0 0]
	fmt.Println("切片长度", len(slice))
	fmt.Println("切片的容量", cap(slice))
	slice[0] = 66
	slice[1] = 77
	slice[2] = 88
	slice[3] = 99
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v \t", i, slice[i])
	}
	hengan.Hengan()
	for key, value := range slice {
		fmt.Printf("下标：%v ,元素： %v \n", key, value)
	}
	hengan.Hengan()
	//底层原理： append 追加
	//1.底层追加元素的时候对数组进行扩容，老数组扩容为新数组
	//2.创建一个新数组，将老数组中的4，7，3复制到新数组中，在新数组中追加80，90
	//3.slice3 底层数组的指向 指向的是新数组
	arr := []int{78, 45, 65, 54, 12}
	var slice2 []int = arr[1:4]
	fmt.Println(len(slice2))
	fmt.Println(slice2)
	slice3 := append(slice2, 80, 90)
	fmt.Println(slice3)
	slice4 := []int{1, 2}
	slice3 = append(slice3, slice4...)
	fmt.Println(slice3)
	hengan.Hengan()
	//切片拷贝
	var a []int = []int{1, 3, 5, 6, 7, 8}
	var b []int = make([]int, 10)
	copy(b, a) //将a中对应的数组元素复制到b中
	fmt.Println(b)
}
