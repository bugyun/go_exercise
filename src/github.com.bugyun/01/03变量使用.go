package main //必须有一个main包

import "fmt" //导入包含，必须使用

func main() {
	//变量，程序运行期间，可以改变的量
	//1.声明格式 var 变量名 类型，变量声明了，必须要使用
	//2.只是声明没有初始化的变量，默认值为0
	//3.同一个{}里，声明的变量名是唯一的
	var a int //默认值为0
	fmt.Println(a)
	a = 10

	//4.可以同时声明多个变量
	var b, c int
	fmt.Println(b, c)

	//3.自动推导类型，必须初始化，通过初始化的值确定类型
	d := "nihao"
	fmt.Println(d)

}
