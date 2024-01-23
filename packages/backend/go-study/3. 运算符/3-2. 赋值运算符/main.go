package main

import "fmt"

func main() {
	// 交换两个int值
	a := 1
	b := 2
	var c int
	fmt.Println(a, b)

	c = a
	a = b
	b = c
	fmt.Println(a, b)
}
