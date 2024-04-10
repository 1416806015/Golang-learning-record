package main

import "fmt"

//定义动物的结构体
type Animal struct {
	Age    int
	Weight float32
}

//给Animal绑定方法：喊叫：
func (an *Animal) Shout() {
	fmt.Println("我可以大声喊叫")
}

//给Animal绑定方法：自我展示：
func (an *Animal) showInfo() {
	fmt.Printf("动物的年龄是：%v ，动物的体重：%v", an.Age, an.Weight)
}

//定义结构体：Cat
type Cat struct {
	//为了复用性：体现继承思维，加入匿名结构体：
	Animal
}

//对Cat绑定特有的方法：
func (c *Cat) scrath() {
	fmt.Println("\n我是小猫，我可以挠人")
}

type Dog struct {
	Animal
}

//对Cat绑定特有的方法：
func (d *Dog) see() {
	fmt.Println("\n我是小狗，我可以看家")
}
func main() {
	//创建Cat结构体示例
	cat := &Cat{}
	cat.Animal.Age = 3
	cat.Animal.Weight = 20.1
	cat.Animal.Shout()
	cat.showInfo()
	cat.scrath()
	//unknown field Age in struct literal of type Dog
	// dog := &Dog{Age: 4, Weight: 20.5} 这是错误的写法
	/*原因：
	Dog 结构体内部嵌入了 Animal 结构体，
	因此 Dog 结构体实例化时不能直接初始化 Age 和 Weight 字段，
	因为它们属于 Animal 结构体。要初始化 Dog 结构体中的 Age 字段，
	需要明确指定 Animal 结构体字段：
	*/
	dog := &Dog{Animal: Animal{Age: 4, Weight: 20.5}}
	dog.Shout()
	dog.showInfo()
	dog.see()
}
