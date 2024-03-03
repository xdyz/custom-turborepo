package main

import (
	"fmt"
	"github.com/xdyz/rabbitmq-demo/internal"
	"time"
)

func main() {
	conn, err := internal.ConnectRabbitMQ("my-mq", "my-mq", "localhost:5672", "")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// create queue for testing

	if err := client.CreateQueue("my-queue", true, false); err != nil {
		panic(err)
	}

	if err := client.CreateQueue("my-queue1", false, true); err != nil {
		panic(err)
	}

	fmt.Print(conn)

	time.Sleep(10 * time.Second)
}
