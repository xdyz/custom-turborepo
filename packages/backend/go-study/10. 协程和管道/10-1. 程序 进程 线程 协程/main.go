package main

import (
	"fmt"
	"time"
)

// main 为主协程，如果主协程执行完毕，主协程内部通过go开启的从协程，将也会关闭，不再执行
func main() {

	// 启动一个协程

	// 协程的执行是异步的，下面两个将会同时执行
	go sayHello()

	// sayMsb 是同步的，这里的代码执行完，主协程main 就结束了，所以 sayHello 剩下的循环将不会在执行
	sayMsb()
}

func sayMsb() {
	for i := 0; i < 3; i++ {
		fmt.Println("msb", i)
		time.Sleep(time.Second) // 睡眠一秒
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello", i)
		time.Sleep(time.Second) // 睡眠一秒
	}
}
