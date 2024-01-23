package main

import "fmt"

func main() {

	fmt.Println("请输入一个整数")
	var a string
	fmt.Scanln(&a)
	fmt.Printf("你输入的整数是：%s\n", a)
}
