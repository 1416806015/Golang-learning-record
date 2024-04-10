package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

/*如果一个类型实现了String（）这个方法，那么fmt.Println默认会调用这个变量
  的String（）进行输出以后定义结构体的话，常定义String（）作为输出结构体信息的方法，fmt.Ptirntln会自动调用*/
func (s *Student) String() string {
	str := fmt.Sprintf("Name = %v , Age = %v", s.Name, s.Age)
	return str
}

func main() {
	stu := Student{
		Name: "老六",
		Age:  20,
	}
	fmt.Println(&stu)
}
