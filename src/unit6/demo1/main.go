package main

import (
	"errors"
	"fmt"
)

// defer+recover机制处理错误
func main() {
	test()
	fmt.Println("-----------------------")
	err := test2()
	if err != nil {
		fmt.Println("自定义错误:", err)
		//停止当前程序的正常执行
		panic(err)
	}
	fmt.Println("上面的除法操作执行成功。。。。")
	fmt.Println("正常执行下面的逻辑。。。。")
}

func test() {
	//利用defer+recover来捕获错误：defer后加上匿名函数的调用
	defer func() {
		//调用recover内置函数，可以捕获错误：
		err := recover()
		//如果没有捕获错误，返回值为零值：nil
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err是：", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

// 自定义错误
func test2() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误：
		return errors.New("除数不能为0哦~~")
	} else { //如果除数不为0，那么正常执行就可以了
		result := num1 / num2
		fmt.Println(result)
		//如果没有错误，返回零值
		return nil
	}

}
