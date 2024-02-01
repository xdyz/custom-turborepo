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

func main() {
	wg.Add(2)
	go writeIndex()
	go panicHandler()
	wg.Wait()

	fmt.Println("main end")

}
