package main

import "fmt"

/**
	交换2个值：
	1.普通方法
	2.go中特有的方法
 */

func normalChange() {
	a, b := 10, 20
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d,b = %d\n", a, b)
}

func goChange() {
	//交换2个变量的值
	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d,j= %d\n", i, j)
}

func main() {
	normalChange()
	goChange()

	i, j := 10, 20
	//_匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势
	tmp, _ := i, j
	fmt.Println("tmp = ", tmp)

}
