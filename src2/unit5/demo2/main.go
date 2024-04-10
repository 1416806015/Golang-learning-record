package main

import (
	"fmt"
	"time"
)

func main() {
	//时间函数 time包  Now函数
	now := time.Now()
	fmt.Printf("%v ~~~ 对应的类型为 %T\n", now, now)
	//Now 返回的是一个结构体  类型是 time.Time
	//调用结构体中的方法
	fmt.Printf(" 年: %v \n ", now.Year())
	fmt.Printf("月: %v \n ", now.Month())      //月：March
	fmt.Printf("月: %v \n ", int(now.Month())) //月：3
	fmt.Printf("日: %v \n ", now.Day())
	fmt.Printf("时: %v \n ", now.Hour())
	fmt.Printf("分: %v \n ", now.Minute())
	fmt.Printf("秒: %v \n ", now.Second())
	fmt.Printf("--------------------------\n")
	//Printf将字符串直接输出
	fmt.Printf("当前的年月日：%d-%d-%d 时分秒：%d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//Sprintf可以得到这个字符串，以便后续使用
	detestr := fmt.Sprintf("当前的年月日：%d-%d-%d 时分秒：%d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(detestr)
	fmt.Printf("--------------------------\n")
	//这个参数字符串的各个数字必须是固定的，必须这样写
	detestr2 := now.Format("2006/01/02 15:04:05")
	fmt.Println(detestr2)
}
