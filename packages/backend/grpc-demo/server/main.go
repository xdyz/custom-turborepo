package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pbService "grpc-demo/service"
	"io"
	"net"
)

// Server 需要把 grpc 中的 UnimplementedDemoServiceServer 作为被继承的结构体，在我们的服务中进行实现
// 所以需要定义一个新的结构体来继承 grpc 生成的 UnimplementedDemoServiceServer 结构体
type Server struct {
	pbService.UnimplementedDemoServiceServer
}

//status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
// 上面的错误会直接返回一个错误给到客户端，无论前面返回了什么数据，此时客户端就收不到返回的的数据了，上面一行的代码优先级比较高

func (s *Server) SayHello(ctx context.Context, req *pbService.User) (*pbService.UserInfo, error) {
	fmt.Println("Hello")
	data := &pbService.UserInfo{
		Name: req.GetUsername(),
		Age:  18,
		Roles: []*pbService.Roles{
			&pbService.Roles{
				Roles: []string{"admin", "user"},
			},
		},
	}

	fmt.Println(data)

	return data, nil
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

	return data, nil
}

// UpdateProductClient server implement the stream key word  method
func (s *Server) UpdateProductClient(stream pbService.ProductService_UpdateProductClientServer) error {
	// stream may be will always from the client to the server, so we need to get it by [for{}]
	for {
		recv, err := stream.Recv()
		if err != nil {
			// when we receive the EOF, it means stream is over
			if err == io.EOF {
				return nil
			}
			return err
		}

		err1 := stream.SendAndClose(recv)
		if err1 != nil {
			return err1
		}
	}
}

// Auth achieve auth method for verify the token
func Auth(ctx context.Context) error {
	mb, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("no metadata")
	}

	var user string
	var password string

	if val, ok := mb["user"]; ok {
		user = val[0]
	}

	if val, ok := mb["password"]; ok {
		password = val[0]
	}

	// if the password and user is not equal , return false
	if user == "admin" && password == "123456" {
		return nil
	}

	return nil
}

// AuthIntercept define an authentication interceptor method to verify each grpc request
func AuthIntercept(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// intercept the request method and verify that token

	err = Auth(ctx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func main() {
	// 添加证书
	// 创建一个自定义证书的服务创建, 需要先用openssl生成证书
	//creds, err := credentials.NewClientTLSFromFile("cert/server.pem", "cert/server.key")

	// 1. new 一个grpc服务端， 这个服务端是不需要认证的
	rpcServer := grpc.NewServer()

	// 这是一个需要认证的服务
	//rpcServer := grpc.NewServer(grpc.Creds(creds))

	// This is a server with interceptor, interceptor for the authentication
	//rpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(AuthIntercept))

	// 2. 将服务注册到grpc上 这个服务在 grpc.pb.go中已经实现了，直接拿过来调用即可
	pbService.RegisterDemoServiceServer(rpcServer, &Server{})

	fmt.Println("Starting")

	// 3. 建立一个当前服务的端口号
	listen, _ := net.Listen("tcp", "127.0.0.1:9090")

	// 4. 调用rpcServer.Serve()方法来启动服务
	err := rpcServer.Serve(listen)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("服务启动成功")
	}

}
