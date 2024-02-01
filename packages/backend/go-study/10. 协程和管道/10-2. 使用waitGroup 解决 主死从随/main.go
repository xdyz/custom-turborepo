package main

import (
	"fmt"
	"sync"
	"time"
)

var wp sync.WaitGroup

func main() {

	// wp.Add(10) // 如果已经知道会启动几个协程，可以提前设置好计数器，不用每次启动一个然后再+1

	for i := 0; i < 10; i++ {
		wp.Add(1) // 协程计数+1
		go worker(i)
	}

	wp.Wait() // 阻塞主线程，什么时候 wp的计数器变为0了，就不阻塞了

	fmt.Println("main") //  前面是阻塞的，只有前面的协程执行完了，这里才会继续执行

}

func worker(n int) {
	defer wp.Done()         // wp.Add(-1)
	time.Sleep(time.Second) // 为了看起来更清晰，这里加了延时，可以不用加
	fmt.Println("worker ", n)

}
