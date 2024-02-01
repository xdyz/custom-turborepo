package main

import (
	"fmt"
	"time"
)

var c = make(chan int, 1)
var s = make(chan string, 1)

func main() {
	go func() {
		c <- 100
		s <- "hello"

		close(c)
		close(s)
	}()

	// 开启一个定时循环执行 一秒执行一次,
	for {
		select {
		case r := <-c:
			fmt.Printf("c: %v\n", r)
		case r := <-s:
			fmt.Printf("s: %v\n", r)
		default:
			fmt.Println("default")
		}

		time.Sleep(time.Second)
	}
}
