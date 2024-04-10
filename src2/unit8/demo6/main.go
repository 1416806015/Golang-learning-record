package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	//1.按顺序赋值
	var s1 Student = Student{"小明", 19}
	fmt.Println(s1)
	//2.按指定类型
	var s2 Student = Student{
		Name: "老六",
		Age:  20,
	}
	fmt.Println(s2)
	//3.想要发货结构体的指针类型
	var s3 *Student = &Student{"明明", 20}
	fmt.Println(*s3)
	var s4 *Student = &Student{
		Name: "老六",
		Age:  20,
	}
	fmt.Println(*s4)
}
