package main

import "fmt"

func main() {
	// 定义一个管道
	var c chan int = make(chan int, 3) // 定义一个容量 为 3的，存储int类型的管道

	// 给管道写入值, 因为管道容量为3，所以最多写入3，如果不取出，继续写入就会报错
	c <- 100
	c <- 200
	c <- 300
	fmt.Printf("管道的实际长度 %v 管道的容量 %v \n", len(c), cap(c))

	// 从管道取出值
	data := <-c
	fmt.Println(data)                                    // 这里为100， 因为第一次存入的先被取出，管道里面还剩下一个值200
	fmt.Printf("管道的实际长度 %v 管道的容量 %v \n", len(c), cap(c)) // 再次打印管道信息，此时实际长度变为1， 容量还是3

	close(c) // 关闭管道 但是，只是关闭管道的进门口，出门口依然可以继续从管道中读取没有取出的数据

	// c <- 600 // 此时会报错，因为管道的进门口已经关闭了，不能再往管道中写数据了
	// fmt.Println(<-c) // 但是还可以将未读取完的数据继续读取, 打印200

	// 只有管道关闭以后，才可以遍历管道，不然会报错
	for v := range c {
		fmt.Println(v) // 直接会遍历出管道中的值
	}

	// ------------------------------------------

	// 这是只写管道，只能往里面写，不能读取数据
	var c2 chan<- int = make(chan<- int, 10)

	// fmt.Println(<-c2) // 直接提示错误
	c2 <- 100 // 但是写入是可以的

	// 定义一个只读管道, 不能写入，只能读取
	var c3 <-chan int = make(<-chan int, 10)
	// fmt.Println(<-c3) // 此时是可以读取的

	// c3 <- 100 // 这个是错误的，因为只读管道不能写入数据

	// 判断一个管道是否为空
	if c3 != nil {
		fmt.Println("c3 不是 nil") // 此时就可以读取了
	}

	var c4 chan int // 没有初始化时，管道是 nil , 空指针

	fmt.Println(c4 == nil) // 此时为true

	var c5 chan int = make(chan int, 10) // 此时初始化了，有一个地址了，所以就不是nil了

	fmt.Println(c5 == nil) // 此时为false
	// c5 = nil // 但是还是可以给赋值空指针的

}
