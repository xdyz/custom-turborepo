//  第一中写法，同时循环，边写边读, 此时才是交替工作

package main

import (
	"fmt"
	"sync"
)

// 定义一个WaitGroup
var wp sync.WaitGroup

// 定义一个管道
var c chan int = make(chan int, 50)

func main() {

	for i := 0; i < 20; i++ {
		go writeData(i)
	}

	for i1 := 0; i1 < 20; i1++ {
		go readData()
	}

	wp.Wait()

	defer close(c)

	fmt.Println("main end")
}

func writeData(n int) {
	fmt.Println("writeData", n)
	wp.Add(1)
	c <- n
}

func readData() {
	fmt.Println("readData ", <-c)
	defer wp.Done()
}

// 第2中写法，等写完，之后在遍历读取管道，但是这种不属于同时在工作

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // 定义一个WaitGroup
// var wp sync.WaitGroup

// func main() {

// 	// 定义一个管道
// 	var c chan int = make(chan int, 50)
// 	wp.Add(2)
// 	go writeData(c)
// 	go readData(c)

// 	wp.Wait()

// 	fmt.Println("main end")
// }

// func writeData(intChan chan int) {
// 	defer wp.Done() // 这个协程执行完了，通知wp减1
// 	for i := 0; i < 20; i++ {
// 		fmt.Println("write data:", i)
// 		intChan <- i
// 	}

// 	close(intChan) // 关闭通道，方便后面读取
// }

// func readData(intChan chan int) {
// 	defer wp.Done() // 这个协程执行完了，通知wp减1
// 	for v := range intChan {
// 		fmt.Println("read data:", v)
// 		fmt.Println(v)
// 	}
// }
