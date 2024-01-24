package main

import "fmt"

func main() {

	// 这种调用方式，就导致 getSum 中的sum 变量一直存在，形成了闭包
	var f = getSum(1)

	res := f(2)

	fmt.Println(res)

	res1 := f(5)

	fmt.Println(res1)
}

func getSum(a int) func(b int) int {
	var sum int = 0
	return func(b int) int {
		sum += a + b
		return sum
	}
}
