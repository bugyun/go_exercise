package main

import "fmt"

var (
	a int
	b float64
)

const (
	i int    = 10
	s string = "nihao"
)

func main() {
	a, b = 19, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	fmt.Println("i =", i)
	fmt.Println("s =", s)

}
