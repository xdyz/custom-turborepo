// package demo2
package main

import "fmt"

// 变量的四种使用方式

var h1 = 100 // 全局变量

// h44:= 200 // 全局不能使用这种方式

func main() {

	// 局部变量

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

	// --------------------------------------------------

	// 同时声明多变量

	var n1, n2, n3 int // 此时 n1, n2, n3 的类型都是 int 都为默认值0

	var n4, n5, n6 int = 10, 20, 30 // 同下
	var (
		a1 int = 10
		b2 int = 20
		c3 int = 30
	)

	// 或者不指定类型 但是必须赋值，不然无法推断类型
	var n7, s1, d1 = 10, "hello", 30.0

	fmt.Println(n1, n2, n3, n4, n5, n6, a1, b2, c3)
	fmt.Println(n7, s1, d1)
}
