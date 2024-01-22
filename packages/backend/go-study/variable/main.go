// package demo2
package main

import "fmt"

// 变量的四种使用方式
func main() {
	// 第一种，定义和赋值
	var a int = 10
	fmt.Println(a)

	// 第二种 只定义，不赋值
	var b int // 此时默认值为 0
	fmt.Println(b)

	var s string // 默认值为 ""
	fmt.Println(s)

	// 第三种，不指定类型，Go会根据赋值的值来推断变量的类型
	var c = 20
	fmt.Println(c)

	// 第四种 省略var关键字 使用 := 来表示一个变量 + 赋值
	d := 30
	fmt.Println(d)
}
