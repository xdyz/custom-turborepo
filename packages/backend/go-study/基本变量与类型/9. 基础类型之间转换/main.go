package main

import "fmt"

func main() {
	var a int = 10
	var b float32 = float32(a) // 强制转换

	fmt.Printf("%T %v \n", a, a) // %T %v 一个打印变量的类型和值  int 10
	fmt.Printf("%T %v", b, b)    // float32 10

	// 精度丢失了
	var c int64 = 87888890
	var d int8 = int8(c)

	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)

	var e int32 = 30
	// var f int64 = e + 40 // 不允许 e + 40 加完是 int32 类型 与 左侧不匹配
	var f int64 = int64(e) + 40 // 可以
	fmt.Printf("%T %v \n", e, e)
	fmt.Printf("%T %v \n", f, f)

	var g int8 = int8(e) + 127 // 不会编译出错，但是溢出导致结果为负数了
	fmt.Printf("%T %v \n", g, g)

	fmt.Sprintf("%T %v \n", g, g)

}
