package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 文件创建/写入
	// fileCreate()

	// 文件追加内容
	// fileAppend()

	// 带缓冲区的写入
	writeWithBuffer()

}

func fileCreate() {
	//os.O_RDWR|os.O_CREATE 读写 和如果不存在这个文件就先创建后写入 如果存在就直接写入
	file, err := os.OpenFile("packages/files/a.txt", os.O_RDWR|os.O_CREATE, 0644)

	// 操作完成后，自动关闭文件
	defer file.Close()

	// 文件打开失败
	if err != nil {
		fmt.Println(err)
	}

	// 按字节写入
	// res, err1 := file.Write([]byte{'a'})

	// 写入字符串，会从头开始写入，如果文件已经存在，会覆盖掉文件内容
	res, err1 := file.WriteString("afafdaf123")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(res)
}

func fileAppend() {
	file, err := os.OpenFile("packages/files/a.txt", os.O_RDWR|os.O_APPEND, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	res, err1 := file.WriteString("afafdaf123")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(res)
}

// 带缓冲区的写入

func writeWithBuffer() {

	//
	file, err := os.OpenFile("packages/files/b.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(file)

	// 此时写入的数据，还没有写入到文中，只是写入到了缓冲区中
	res, err := writer.WriteString("a额2噢 daf123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// 将缓冲区的数据写入到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

	// 修改文件权限
	// os.FileMode(0644).String()

}
