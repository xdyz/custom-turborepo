package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间, Now 方法 返回的是一个结构体，可以当做类看待
	now := time.Now()

	fmt.Println(now)

	// 与 JS 不同 如果要获取 YYYY-MM-DD hh:mm:ss 的格式，需要使用 Format 方法
	// 必须使用 2006-01-02 15:04:05 的字符串传入，这个是go的默认

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	time := fmt.Sprintln(now.Format("2006-01-02 15:04:05"))
	fmt.Println(time)
}
