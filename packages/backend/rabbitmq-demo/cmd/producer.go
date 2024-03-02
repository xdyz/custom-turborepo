package cmd

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	// connect to amqp (rabbitmq), use docker to start a rabbitmq sever and a rabbitmq management sever
	conn, err := amqp.Dial("localhost")

	if err != nil {
		panic(err)
	}

	defer conn.Close()
}
