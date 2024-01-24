package main

func main() {
	// 定义匿名函数，直接调用。自执行函数
	func(a int, b int) int {
		return a + b
	}(1, 2)

	// 定义匿名函数，赋给一个变量, 最后调用

	result := func(a int, b int) int {
		return a + b
	}
	result(1, 2)
}
