package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func writeIndex() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func panicHandler() {
	defer wg.Done()
	// 如果不想让程序退出，可以使用recover, 错会输出，但是程序不会崩溃，会继续执行下去，而不是直接退出
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover err:", err)
		}
	}()
	num := 10
	num1 := 0

	num2 := num / num1 // 此时会报错，直接panic 程序退出了

	fmt.Println(num2)

}

type D struct {
	Name string
}

func main() {

	var a map[string]int
	fmt.Println(a == nil)

	var b []int
	fmt.Println(b == nil)
	// var c [3]int
	// fmt.Println(c == nil)
	// var d D
	// fmt.Println(d == nil)

	// 结构体声明一个 变量
	var d D
	fmt.Println("%v", d)

	// wg.Add(2)
	// go writeIndex()
	// go panicHandler()
	// wg.Wait()

	// fmt.Println("main end")

}
