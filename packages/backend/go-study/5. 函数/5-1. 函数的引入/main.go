package main

import "fmt"

func funcName() (int, string) {

	// 函数体
	return 1, "hello"

}
func main() {
	a, b := funcName()

	fmt.Println(a, b)
}
