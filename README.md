# Golang learning record

# unit

## demo 01-02

变量与数据类型

## demo 03

类型转换，二进制溢出

## demo 04

基本型转string类型 

1.fmt.Sprintf("%参数",表达式)

2.使用strconv包函数    

## demo 05

string类型转基本型

    strconv函数 

## demo 06

指针

         1.  & 取地址 2.  * 取值

# unit2

## demo 1

出现的问题：

    引入自己创建的包，运行时出现编译器无法找到它的位置。即使配置了环境变量，编译器仍然无法正确找到包的位置。这是由于从Go的1.11版本之后，已不再推荐使用GOPATH来构建应用了

![](C:\Users\chen\AppData\Roaming\marktext\images\2024-03-04-15-27-10-image.png)

解决方案：

```
go env  //查看当前操作系统环境中的所有环境变量

将GO111MODULE修改为off，在终端输入
go env -w GO111MODULE=off
```

### GO111MODULE

有三个可选值：`off`、`on`、`auto`，默认值是 `auto`。

GO111MODULE=off 无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
GO111MODULE=on 模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod 下载依赖。
GO111MODULE=auto 在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持。


## 变量命名

如果变量名，函数名，常量名首字母大写，则可以被其他的包访问；

如果首字母小写，则只能在本包中使用（利用首字母的大小写完成权限控制）
