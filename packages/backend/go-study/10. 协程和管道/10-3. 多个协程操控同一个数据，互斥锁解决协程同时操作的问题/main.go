package main

import (
	"fmt"
	"sync"
)

var total int

var mutex sync.Mutex //  互斥锁
var wp sync.WaitGroup

func main() {
	wp.Add(2)
	go addNum()
	go deNum()

	wp.Wait()

	fmt.Println(total) // 这个结果并不是预期的 0 ，而是其余的数字

}

func addNum() {
	defer wp.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()   // 使用之前先加锁
		total += 1     // 中间部分就被锁起来了
		mutex.Unlock() // 使用完成之后就解锁
	}
}

func deNum() {
	defer wp.Done()
	for i := 0; i < 10000; i++ {
		mutex.Lock()   // 使用之前先加锁
		total -= 1     // 中间部分就被锁起来了
		mutex.Unlock() // 使用完成之后就解锁
	}
}
