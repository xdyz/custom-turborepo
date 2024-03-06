package main

import (
	"fmt"
	"github.com/xdyz/rabbitmq-demo/internal"
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

	// create a consumer, only used for get messages( <- chan amqp.Delivery )
	messageBus, err := client.Consume("my-queue", "my-consume", false, nil)

	if err != nil {
		panic(err)
	}

	var blocking chan struct{}

	// 启动协程，来一直接收消息
	go func() {
		for message := range messageBus {
			fmt.Println(message)

			// 判断此消息是不是重新投递过
			if !message.Redelivered {
				// Nack 方法通知rabbitmq，此消息已经被消费了，但是没有确认消费成功，所以rabbitmq会重新投递
				// 第一个参数表示是否应该同时拒绝当前消息之前队列中的所有未确认消息
				// 第二个参数表示是否应该将拒绝的消息重新放回队列
				message.Nack(false, true)
				continue
			}

			// 传入 false 就是只确认当前消息是否被成功消费了，如果为true, 就是告诉rabbitmq从队列头到当前消息 之间的所有消息都被成功消费
			if err := message.Ack(false); err != nil {
				fmt.Println("message failed", err)
				continue
			}

			fmt.Println("message id", message.MessageId)
			me := message.Body
			fmt.Println("body", me)
		}
	}()

	// 我们没有给这个消息队列发送消息，所以这里会阻塞，上面的协程就会一直运行，main 不退出，协程不退出
	<-blocking
}
