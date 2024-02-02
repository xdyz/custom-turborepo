package main

import (
	"fmt"
	"net"
)

func main() {
	// 传入两个参数，一个是链接类型，一个是ip+port( 服务端所在 ip 和 port)
	conn, err := net.Dial("tcp", "127.0.0.1:8888") // 启动一个客户端, 如果成功就返回一个连接
	if err != nil {
		fmt.Println("客户端连接失败 \n", err)
		return
	} else {
		fmt.Println("客户端连接成功 \n", conn)
	}

	n, err1 := conn.Write([]byte("hello world")) // 向服务端发送数据
	if err1 != nil {
		fmt.Println("客户端发送数据失败 \n", err1)
		return
	} else {
		fmt.Printf("客户端发送数据成功 %v \n", n)
	}

	defer conn.Close()
}
