package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"grpc-demo/service"
)

func main() {
	user := &service.User{
		Password: "password",
		Username: "username",
	}

	// 序列化的过程
	marshal, err := proto.Marshal(user)

	if err != nil {
		fmt.Println(err)
	}

	// 反序列化
	err1 := proto.Unmarshal(marshal, user)
	if err1 != nil {
		fmt.Println(err1)
	}
}
