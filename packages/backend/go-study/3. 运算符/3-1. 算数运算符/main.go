package main

import "fmt"

func main() {
	// 加法
	age := 10 + 8
	fmt.Println(age)

	// 减法
	age1 := 10 - 8
	age2 := 10.3 - 8
	fmt.Println(age1)
	fmt.Println(age2)

	// 除法
	num := 5 / 2      // 两个都为int 类型的相除，得到的只有整数位，不会保留小数位
	num1 := 10 / 2.33 // 一个 float 和 一个int 相除 得到的就是 float 类型

	fmt.Println(num)
	fmt.Println(num1)

	// 乘法
	num2 := 10 * 2.33
	fmt.Println(num2)

	// 自增
	age++ // 不能使用 age3 := age++
	fmt.Println(age)

	// 自减

	age--
	fmt.Println(age)
}
