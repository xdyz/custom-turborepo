package main

import (
	"fmt"
	"sync"
	"time"
)

var total int

var rwMutex sync.RWMutex //  读写锁
var wp sync.WaitGroup

func main() {
	wp.Add(1)

	for i := 0; i < 3; i++ {
		wp.Add(1)
		go read()
	}
	go write()

	wp.Wait()

	fmt.Println(total) // 这个结果并不是预期的 0 ，而是其余的数字

}

func read() {
	defer wp.Done()
	rwMutex.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据完成")
	rwMutex.RUnlock()
}

// 写入的时候，读取的操作会阻塞
func write() {
	defer wp.Done()
	rwMutex.Lock() // 写加锁
	fmt.Println("开始写入数据")
	time.Sleep(time.Second * 5)
	fmt.Println("写入数据完成")
	rwMutex.Unlock()
}
