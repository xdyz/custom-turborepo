package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("start")
	// test() // 此时虽然会panic 但是不会中断程序，会继续执行
	err := test1() // 自定义抛出错误

	// 主动中断程序
	if err != nil {
		panic(err)
	}

	fmt.Println("end")
}

func test() {
	// defer + recover 来捕获错误，并且让函数继续执行，不会直接中断程序
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("错误已捕获")
			fmt.Println("Recovered in f", r)
		}
	}()

	num1 := 10

	num2 := 0

	result := num1 / num2

	fmt.Println(result)
}

// 自定义错误
func test1() (err error) {
	num1 := 10

	num2 := 0

	if num2 == 0 {
		// 自己抛出错误
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)

		return nil
	}

}
