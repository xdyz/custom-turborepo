package main

import "fmt"

func main() {
	fmt.Println("请输入一个整数")
	var a int
	fmt.Scanln(&a) //  传入的是一个地址，因为函数内向改变传入的值
}
