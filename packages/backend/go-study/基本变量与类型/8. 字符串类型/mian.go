package main

import "fmt"

func main() {
	// 1. 定义一个字符串
	var s1 string = "你好"
	fmt.Println(s1)

	// 2. 字符串不可变：不是指给予变量值不可变，指的是其中的字符值不可变
	// s1[0] = 'g' //这个是错误的

	// 用 `` 定变量
	var s22 = `你好啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊啊`
	fmt.Println(s22)

	// 3. 字符串的拼接
	var s2 = "你好"
	var s3 = "啊啊啊"
	var s4 = s2 + s3
	fmt.Println(s4)

}
