package main

import "fmt"

// 函数text 返回值为int类型
func text() int {
	return 10
}

func main() {

	// 单分支
	var a int = 50

	if a > 20 {
		fmt.Println("a > 20")
	}

	if a < 20 {
		fmt.Println("a < 20")
	} else {
		fmt.Println("a >= 20")
	}

	// 多分支
	if a < 20 {
		fmt.Println("a < 20")
	} else if a >= 20 && a < 30 {
		fmt.Println("a >= 20")
	} else if a >= 30 && a < 40 {
		fmt.Println("a >= 30")
	} else if a >= 40 && a < 50 {
		fmt.Println("a >= 40")
	} else {
		fmt.Println("a >= 50")
	}

	// switch

	var b int = 10

	switch b {
	case 10:
		fmt.Println("a = 10")
		fallthrough // 带穿透 此时上下两个case 都会被打印出来
	case 20:
		fmt.Println("a = 20")
	case 30:
		fmt.Println("a = 30")
	case 40:
		fmt.Println("a = 40")
	case 50:
		fmt.Println("a = 50")
	default:
		fmt.Println("a = ?")
	}

	// switch 后面不跟 变量也可以  当 if else 来使用

	switch {
	case a < 20:
		fmt.Println("a < 20")
	case a >= 20 && a < 30:
		fmt.Println()
	}

	// 还可以是函数 带返回值的
	switch text() {
	case 10:
		fmt.Println("a = 10")
	case 20:
		fmt.Println("a = 20")
	case 30:
		fmt.Println("a = 30")
	case 40:
		fmt.Println("a = 40")
	case 50:
		fmt.Println("a = 50")
	default:
	}

	// switch 一个case 可以有多个匹配值或者条件 以逗号分隔
	switch b {
	case 10, 20, 30:
		fmt.Println("a = 10, 20, 30")
	case 40, 50:
		fmt.Println("a = 40, 50")
	default:
		fmt.Println("a = ?")
	}
}
