package main //package进行包的声明 建议：包的声明和这个包所在的文件夹同名

//包名是从$GOPATH/src/后开始计算的
/*
import (
	"fmt"
	"unit4/demo3/crm/dbutils"
)

func main() {
	fmt.Println("执行了main")
	dbutils.GetConn2()
	dbutils.GetConn() //在函数调用的时候前面要定位到所在的包
}
*/

//也可以取别名
import (
	"fmt"
	test "unit4/demo3/crm/dbutils"
)

func main() {
	fmt.Println("执行了main")
	test.GetConn2()
	test.GetConn() //在函数调用的时候前面要定位到所在的包
}
