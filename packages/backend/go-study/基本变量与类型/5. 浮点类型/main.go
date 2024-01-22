package main

import "fmt"

func main() {
	var num1 float32 = 3.14

	var num2 float32 = 314e-2 // 314 * 10^-2  也为3.14

	var num3 float32 = 314e+2 // 314 * 10^2  也为31400 加号可省略

	fmt.Println(num1, num2, num3)

	var num4 float64 = -3.14e-10

	fmt.Println(num4)

	// 精度丢失
	var num5 float32 = 256.000000916 // 显示为256
	var num6 float64 = 256.000000916 // 显示正确

	fmt.Println(num5, num6)

	var num = 3.14

	fmt.Printf("%T", num)

}
