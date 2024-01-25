package main

import "fmt"

func main() {
	fmt.Println(len("daljfj你")) // 返回所占的字节数

	// new 返回一个指向一个类型为T的零值的指针。
	var p *int = new(int)
	n := new(float32)

	fmt.Println(p, n)
}
