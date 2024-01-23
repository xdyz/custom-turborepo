package main

import "fmt"

func main() {
	// 1 &
	var a, b int = 12, 25 // 12的二进制表示为1100，25的二进制表示为11001
	result1 := a & b      // 结果为8，二进制表示为1000

	fmt.Println(result1)

	// 2 |
	result2 := a | b // 结果为29，二进制表示为11101
	fmt.Println(result2)

	// 3 ^ 按位异或
	result3 := a ^ b // 结果为21，二进制表示为10101
	fmt.Println(result3)

	// 4 ^ 按位取反
	result4 := ^a // 结果为-13，二进制表示为11111111111111111111111111110011（假设int类型为32位）
	fmt.Println(result4)

	// << 左移
	result5 := a << 2 // 结果为48，二进制表示为110000
	fmt.Println(result5)

	// >> 右移
	result6 := a >> 2 // 结果为3，二进制表示为11
	fmt.Println(result6)
}
