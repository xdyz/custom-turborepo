package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. string 转为 bool类型

	var s string = "aafff"

	b, _ := strconv.ParseBool(s) // 这个函数返回的是两个值 所以接收的时候也要按两个接收，_ 关键字 代表不关心这个返回值 忽略

	fmt.Printf("%T, %v \n", b, b) // false

	// 2. string 转为 int 类型
	var s1 string = "123f"

	c, _ := strconv.ParseInt(s1, 10, 64) // 函数接收三个值，第一个参数为要转换的字符串，第二个为转换的进制，第三个为转换的类型

	fmt.Printf("%T, %v \n", c, c) // 123f

	// string 转为 float 类型
	var s2 string = "123.456"
	d, _ := strconv.ParseFloat(s2, 64) // 接收两个值，一个为要转换的字符串，一个就是要转的结果类型，只有float32 和 float64
	fmt.Printf("%T, %v", d, d)

	var s3 string = "goalng"
	f, _ := strconv.ParseFloat(s3, 64) // 接收两个值，一个为要转换的字符串，一个就是要转的结果类型，只有float32 和 float64
	fmt.Printf("%T, %v", f, f)         // golang 转 float  为无效值， 所以就是默认值0  ”123E-2“ 为有效值 可以正确转换 ”123“ 同样
}
