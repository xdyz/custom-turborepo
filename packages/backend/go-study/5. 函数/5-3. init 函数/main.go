package main

import "fmt"

var Num int = test() // 最先执行

func test() int {
	fmt.Print("test\n")
	return 10
}

func main() {
	fmt.Println("hello") // 第三个执行
}

// 先于main()执行
func init() {
	fmt.Print("init\n") // 第二个执行
}

func init() {
	fmt.Print("init2\n") // 第二个执行
}
