package main

import (
	"fmt"
	"unit8/demo10/model"
)

func main() {
	//创建person结构体示例：
	p := model.NewPerson("老六")
	p.SetAge(20)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
}
