package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 调用一次性读取函数
	// oneceRead()

	// 调用带缓冲的读取函数
	readWithBuffer()
}

func oneceRead() {
	// 一次性读取
	// 调用 ReadFile 不需要 Open 和 close 因为  ReadFile 已经封装好了
	// 起始位置 为读取项目根目录
	bytes, err := os.ReadFile("packages/files/a.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 转成字符串，直接显示内容
	fmt.Println(string(bytes))
}

func readWithBuffer() {
	// 带缓冲的读取

	file, err := os.Open("packages/files/a.txt")
	// 关闭文件
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	// 创建一个缓冲区读取器，传入文件结构体
	reader := bufio.NewReader(file) //  接收两个参数，一个是文件，一个是缓冲区大小（默认为4k）

	// 读取流是一个指针类型，其中有读取数据内容的方法 我们通过循环来读取数据，遇到EOF 尾回调就结束
	for {
		line, err := reader.ReadString('\n') // 每次读取一行数据，读到换行符这一次读取就结束

		// 判断这一行的读取 是否报错
		if err != nil {
			fmt.Println("error = ", err)
		}

		fmt.Println(line)

		// 读取到 EOF 尾回调 （内容读完了）就结束
		if err == io.EOF {
			break
		}
	}
}
