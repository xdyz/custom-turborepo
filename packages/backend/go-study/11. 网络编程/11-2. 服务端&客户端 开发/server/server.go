package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("server")

	// 服务器端启动，也是传染 链接类型 和 ip+port
	listen, err := net.Listen("tcp", "127.0.0.1:8888") // 生成一个服务实例
	if err != nil {
		fmt.Println("创建服务器端失败 ", err)
		return
	} else {
		fmt.Println("创建服务器端成功 ", listen)
	}

	// 监听成功之后，等待服务器端连接，不停的循环，看有没有客户端连接过来
	for {
		conn, err := listen.Accept() // 服务实例 启动连接
		if err != nil {
			fmt.Println("接受客户端连接失败 \n", err)
			return
		} else {
			fmt.Printf("接受客户端连接成功 con=%v \n", conn.RemoteAddr().String())
			// 每一次的连接 都准备一个协程，处理客户端发来的消息，处理完就关闭这次连接
		}
		go process(conn)
	}

}

func process(conn net.Conn) {

	defer conn.Close() // 每次连接完成就关闭连接

	for {
		// 读取客户端发来的数据
		data := make([]byte, 1024) //
		n, err := conn.Read(data)  // n 是一个下标  Read 方法 接收一个 []byte 切片参数
		if err != nil && err != io.EOF {
			fmt.Println("读取客户端数据失败 \n", err)
			return
		} else {
			fmt.Printf("读取客户端数据成功 \n")
			fmt.Printf("data: %v\n", string(data[:n])) // 字节切片，读取数据要要转换成字符串 强制转换
			return
		}
	}
}
