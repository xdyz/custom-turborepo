package main

import "fmt"

func main() {

	var b1 byte = 'a'
	fmt.Println(b1) // 97 打印出来是一个数字

	fmt.Printf("%c", b1) // 输出 a 字符 而不是数字 所以fmt 可以用来转换数据

}
