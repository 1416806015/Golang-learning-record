package dbutils

import "fmt"

func GetConn() { //首字母大写，可以被其他包访问
	fmt.Println("执行了dbutils包下的getConn")
}
