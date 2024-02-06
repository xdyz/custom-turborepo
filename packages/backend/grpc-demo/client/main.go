package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pbService "grpc-demo/service"
)

func main() {

	// client also needs cart, otherwise will report an error
	//file, err := credentials.NewClientTLSFromFile("client/keys/mszlu.pem", "*.mszlu.com")
	//if err != nil {
	//	fmt.Println(err)
	//}
	// read the car file
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(file))

	// connect to the server, and use insecure.NewCredentials() method to create an empty Certificate
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	// connect to the grpc client
	client := pbService.NewDemoServiceClient(conn)

	fmt.Println(client)

	// 执行rpc的调用，可以像调用本地方法一样，调用grpc上的方法
	res, _ := client.SayHello(context.Background(), &pbService.User{Username: "test", Password: "123456"})
	fmt.Println(res.GetAge())
	fmt.Println(res.GetName())
	fmt.Println(res.GetRoles())

	res1, _ := client.GetPerson(context.Background(), &pbService.User{Username: "test", Password: "123456"})
	fmt.Println(res1.GetInfo())
}
