package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pbService "grpc-demo/service"
	"net"
)

// Server 需要把 grpc 中的 UnimplementedDemoServiceServer 作为被继承的结构体，在我们的服务中进行实现
// 所以需要定义一个新的结构体来继承 grpc 生成的 UnimplementedDemoServiceServer 结构体
type Server struct {
	pbService.UnimplementedDemoServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *pbService.User) (*pbService.UserInfo, error) {
	fmt.Println("Hello")
	data := &pbService.UserInfo{
		Name: req.Username,
		Age:  18,
		Roles: []*pbService.Roles{
			&pbService.Roles{
				Roles: []string{"admin", "user"},
			},
		},
	}

	fmt.Println(data)

	return data, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (s *Server) GetPerson(ctx context.Context, req *pbService.User) (*pbService.PersonInfo, error) {

	data := &pbService.PersonInfo{
		Info: []*pbService.PersonInfo_Person{
			&pbService.PersonInfo_Person{
				Name: req.Username,
				Age:  33,
			},
			&pbService.PersonInfo_Person{
				Name: "李四",
				Age:  20,
			},
			&pbService.PersonInfo_Person{
				Name: "王五",
				Age:  22,
			},
		},
	}

	return data, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}

func main() {

	// 当前服务的端口号
	listen, _ := net.Listen("tcp", "127.0.0.1:9090")

	// 1. new 一个grpc服务端
	rpcServer := grpc.NewServer()

	// 2. 将服务注册到grpc上 这个服务在 grpc.pb.go中已经实现了，直接拿过来调用即可
	pbService.RegisterDemoServiceServer(rpcServer, &Server{})

	fmt.Println("Starting")

	// 3. 调用rpcServer.Serve()方法来启动服务
	err := rpcServer.Serve(listen)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("服务启动成功")
	}

}
