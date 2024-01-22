package main

import "fmt"

func main() {
	// 变量的声明
	var age int // 可以和下面的合成一句
	// 变量的赋值
	age = 18

	// 变量的声明 + 赋值
	var age1 int = 20

	age2 := 30 // 简写，go 会自动推导出类型

	// 变量的使用
	fmt.Println(age)
	fmt.Println(age1)
	fmt.Println(age2)

	// 变量不可以重复定义，类型与值必须一致
	// var a int = 12.33  error

	// 多个类型一起定义
	var (
		age3 int    = 30
		name string = "张三"
	)

	fmt.Println(age3, name)

}
