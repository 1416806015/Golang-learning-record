package main

import (
	"fmt"
	hengan "unit2/hen"
)

// 映射
// 键值对： 一对匹配的信息
// 语法： var map变量名map[keytype]valuetype
// 特点：1.map集合在使用前一定要make 2.key-value可以无序 3. key不可以重复，否则后一个覆盖前一个
func main() {
	//定义map变量
	/*var a map[int]string
	//只声明map内存是没有分配空间
	//必须通过make函数进行初始化，才会分配空间
	a = make(map[int]string, 10) //map可以存放10个键值对*/
	var a = make(map[int]string)
	//将键值对存入map
	a[200011128] = "陈洪杰"
	a[200011129] = "张三"
	a[200011130] = "小米"
	fmt.Println(a)

	c := map[int]string{
		202011128: "陈洪杰",
		202011129: "小明",
	}
	c[202022233] = "李华"
	//删除  delete(map,key)
	delete(c, 202011128)
	fmt.Println(c)
	//查找 value,bool = map[key]
	value, flag := c[202022233]
	fmt.Println(value)
	fmt.Println(flag)
	hengan.Hengan()
	//map只支持 for-range
	for key, value := range c {
		fmt.Printf("key:%v value:%v \t", key, value)
	}
	hengan.Hengan()
	d := make(map[string]map[int]string)
	d["班级1"] = make(map[int]string, 4)
	d["班级1"][202011100] = "周志"
	d["班级1"][202011101] = "搜索"
	d["班级1"][202011102] = "所属"
	d["班级1"][202011103] = "实数"
	d["班级2"] = make(map[int]string, 4)
	d["班级2"][202011110] = "商务"
	d["班级2"][202011111] = "首位"
	d["班级2"][202011112] = "亲切"
	d["班级2"][202011113] = "手误"
	// 这个写法会有警告: simplify range expression  可以简化写为下面的
	// for k1, v1 := range d {
	// 	fmt.Println(k1)
	// 	for k2, _ := range v1 {
	// 		fmt.Printf("学生学号为：%v\t", k2)
	// 	}
	// 	fmt.Println()
	// }
	for k1 := range d {
		fmt.Println(k1)
		for k2 := range d[k1] {
			fmt.Printf("学生学号为：%v\t", k2)
		}
		fmt.Println()
	}
}
