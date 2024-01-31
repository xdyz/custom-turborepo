package main

import (
	"fmt"
	"os"
)

func main() {

	// 打开文件
	// 返回两个参数 一个为 File结构体，包含了对File的各种操作和属性
	// 第二个参数为错误信息
	// 如果文件不存在，会返回错误信息
	file, err := os.Open("./a.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(file.Name()) // 返回文件名

	defer file.Close() // 文件使用完成后，关闭打开的文件

}
