package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "packages/files/a.txt"
	targetPath := "packages/files/b.txt"

	rwFileContent(path, targetPath)
}

// 一边读取文件，一边将内容写入到新文件
func rwFileContent(path string, target string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	targetFile, targetErr := os.OpenFile(target, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	defer file.Close()
	defer targetFile.Close()

	if err != nil {
		panic(err)
	}

	if targetErr != nil {
		panic(targetErr)
	}

	reader := bufio.NewReader(file)

	writer := bufio.NewWriter(targetFile)

	// 一边从文件中读取，一边写入到目标文件中
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			fmt.Println(err)
		}

		if err == io.EOF {
			break
		}

		res, err1 := writer.WriteString(line + "new\n")
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println(res)

		// 每次写完，需要刷新新文件的内容
		err2 := writer.Flush()
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	return
}
