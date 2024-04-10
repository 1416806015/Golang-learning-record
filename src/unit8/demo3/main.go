package main

import "fmt"

type Person struct {
	Name string
}

// 给Person结构体绑定方法test
func (p Person) test() {
	fmt.Println(p.Name)
}
func main() {
	var p Person
	p.Name = "老六"
	p.test()

}
