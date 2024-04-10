package main

import "fmt"

type Student struct {
	Age int
}
type Person struct {
	Age int
}
type Stu Student

func main() {

	var s Student = Student{10}
	var p Person = Person{10}
	//1、结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)
	//2、结构体进行type重新定义（相当于取别名）Golang认为是新的数据类型，但是相互间可以强转
	var s1 Student = Student{19}
	var s2 Stu = Stu{12}
	s1 = Student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
