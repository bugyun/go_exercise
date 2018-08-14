package main

import "fmt"

/**
	数据类型
		作用:告诉编译器这个数（变量）应该以多大的内存存储
*/

var a1 int

func main() {
	a1 = 12
	a2 := 12
	fmt.Println(a1, a2)
}
