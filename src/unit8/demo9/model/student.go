package model

/* 第一种  大写，可以被直接调用
type Student struct {
	Name string
	Age  int
}*/

//第二种方法 小写  不可直接被外界访问
type student struct {
	Name string
	Age  int
}

//工厂模式
func NewStudent(n string, a int) *student {
	return &student{n, a}
}
