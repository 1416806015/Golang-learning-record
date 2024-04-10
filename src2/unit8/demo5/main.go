package main

import "fmt"

//方法和函数的区别：
/*1.绑定指定类型
	方法：需要绑定指定数据类型
	函数：不需要绑定数据类型
  2.调用方式不一样：
	方法调用方式：函数名(实参列表)
	函数调用方式：变量.方法名(实参列表)
*/
type Student struct {
	Name string
}

// 定义方法：
func (s Student) test01() {
	fmt.Println(s.Name)
}

//定义函数：
func method01(s Student) {
	fmt.Println(s.Name)
}
func main() {
	//调用函数：
	var s Student = Student{"老六"}
	method01(s)
	//调用方法:
	s.test01()
}
