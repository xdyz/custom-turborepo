package internal

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitmqClient struct
type RabbitmqClient struct {
	// The connection used by client
	conn *amqp.Connection
	// Channel is used to process / Send messages
	ch *amqp.Channel
}

// ConnectRabbitMQ connect to RabbitMQ server
func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// NewRabbitMQClient create a client connection
func NewRabbitMQClient(conn *amqp.Connection) (RabbitmqClient, error) {
	// create a channel
	ch, err := conn.Channel()
	if err != nil {
		return RabbitmqClient{}, err
	}

	rc := RabbitmqClient{
		conn: conn,
		ch:   ch,
	}

	return rc, nil
}

// Close the rabbitmq connection
func (r RabbitmqClient) Close() error {

	// Only close the channel, if someone use it
	return r.ch.Close()
}

// CreateQueue  create queue by channel
func (r RabbitmqClient) CreateQueue(queueName string, druable bool, autoDelete bool) error {
	// create queue by channel
	_, err := r.ch.QueueDeclare(
		queueName,  // 队列的名称
		druable,    // 队列是否持久化
		autoDelete, // 是否自动删除，例如 为true, rabbitmq 重启后，队列被自动删除
		false,      // 独占表示，表示只有创建他的链接可以使用，别的不可以
		false,      // 不等待
		nil,        // 其他参数
	)

	return err
}

func (r RabbitmqClient) CreateExchange(exchangeName string, kind string) error {
	return r.ch.ExchangeDeclare(
		exchangeName,
		kind,
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // arguments
	)
}

func (r RabbitmqClient) CreateBinding(queueName, routerKey, exchange string) error {
	return r.ch.QueueBind(
		queueName, // 队列名称
		routerKey, // 路由key值
		exchange,  // 交换机名称
		false,
		nil)
}
