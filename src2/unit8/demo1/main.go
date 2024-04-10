package main

import (
	"fmt"
	hengan "unit2/hen"
)

type Teacher struct {
	Name   string
	Age    int
	School string
}

func main() {
	//1
	var t1 Teacher //未赋值时默认：{ 0 }
	fmt.Println(t1)
	t1.Name = "马士兵"
	t1.Age = 45
	t1.School = "清华大学"
	fmt.Println(t1)
	fmt.Println(t1.Age + 10)

	hengan.Hengan()
	//2创建老师结构体的实例、对象、变量
	var t2 Teacher = Teacher{"马士兵", 45, "清华大学"}
	fmt.Println(t2)
	hengan.Hengan()

	//3
	var t3 *Teacher = new(Teacher) //t是指针，t其实指向的就是地址，应该给这个地址的指向的对象的字段赋值
	(*t3).Name = "马士兵"             // * 的作用：根据地址取值
	(*t3).Age = 45                 //可以简写为下面的赋值方式
	t3.School = "清华大学"
	fmt.Println(t3)

	//4
	var t4 *Teacher = &Teacher{"马士兵", 45, "清华大学"}
	fmt.Println(t4)
}
