package main

import (
	"fmt"
	model "unit8/demo9/model"
)

/*
func main() {
	//挎包创建结构体Student的实例
	//var s model.Student = model.Student{"明明", 20}
	//s := model.Student{"明明", 20}  这样写有错误 unit8/demo9/model.Student struct literal uses unkeyed fields
	//因为在创建 Student 结构体的实例时，你没有使用字段名，而是按照结构体中字段的顺序提供了值。为了解决这个错误，应该使用字段名来初始化结构体实例。
	s := model.Student{Name: "明明", Age: 20}
	fmt.Println(s)
}*/
//小写挎包调用
func main() {
	s := model.NewStudent("小明", 20)
	fmt.Println(*s)
}
