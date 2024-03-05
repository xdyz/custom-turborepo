package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
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

	// create the exchange
	if err := client.CreateExchange("my-exchange", "topic"); err != nil {
		panic(err)
	}

	// bind the queue to the exchange
	if err := client.CreateBinding("my-queue", "my.queue.exchange", "my-exchange"); err != nil {
		panic(err)
	}
	if err := client.CreateBinding("my-queue1", "my.queue.exchange1", "my-exchange"); err != nil {
		panic(err)
	}

	// 创建一个 10s 的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// 发送消息  持久化的消息
	if err := client.Send(ctx, "my-exchange", "my.queue.exchange", amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent, // 持久化的消息，没有消费就会一直存在，前提，队列也是持久化的
		Body:         []byte("first message"),
	}); err != nil {
		panic(err)
	}

	// 发送消息  短暂的消息， 重启会消失
	if err := client.Send(ctx, "my-exchange", "my.queue.exchange1", amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Transient, // 重启rabbitmq 后就丢失了
		Body:         []byte("second message"),
	}); err != nil {
		panic(err)
	}

	fmt.Print(conn)

	time.Sleep(30 * time.Second)
}
