package main

import "fmt"

func main() {
	a := 10
	fmt.Println("a =", a) //一段一段处理，自动加换行

	//格式化输出，把a的内容放在%d的位置
	//"a= 10\n" 这个字符串输出到屏幕，"\n"代表换行符
	fmt.Printf("a = %d\n", a)




}
