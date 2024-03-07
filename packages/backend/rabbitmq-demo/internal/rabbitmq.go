package internal

import (
	"context"
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
func (r RabbitmqClient) CreateQueue(queueName string, druable bool, autoDelete bool) (amqp.Queue, error) {
	// create queue by channel
	// 如果创建的时候，没有传递队列名称，rabbitmq 会自动创建一个队列名称
	q, err := r.ch.QueueDeclare(
		queueName,  // 队列的名称
		druable,    // 队列是否持久化
		autoDelete, // 是否自动删除，例如 为true, rabbitmq 重启后，队列被自动删除
		false,      // 独占表示，表示只有创建他的链接可以使用，别的不可以
		false,      // 不等待
		nil,        // 其他参数
	)

	/**
	 * amqp.Queue{}
	 * @param Name string 队列名称
	 * @param Messages int 未等待确认的消息数量
	 * @param Consumers int 消费者数量
	 */
	if err != nil {
		return amqp.Queue{}, err
	}

	return q, nil
}

// CreateExchange create an exchange , set the echanageName and kind ( or type : topic, fanout, ...etc)
func (r RabbitmqClient) CreateExchange(exchangeName string, kind string) error {
	return r.ch.ExchangeDeclare(
		exchangeName,
		/** 交换机类型
		 * fanout: 扇形、广播给挂在到这个交换机上的所有队列，此时会忽略，所有队列都会收到同一条消息
		 * topic: 通过路由匹配到符合routerKey的队列，给其传递消息 会用到 # 号  *号等
		 * direct: 通过路由找到到与routerKey一致的队列，给其传递消息
		 * headers: 通过headers匹配到符合headers的队列，给其传递消息
		 */
		kind,
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // arguments
	)
}

// CreateBinding let queue bind to the exchange, use routerKey
func (r RabbitmqClient) CreateBinding(queueName, routerKey, exchange string) error {
	return r.ch.QueueBind(
		queueName, // 队列名称
		routerKey, // 路由key值
		exchange,  // 交换机名称
		false,
		nil)
}

// Send Publish 将消息传递给交换机，所以需要交换机的名称，路由键，因为交换机需要知道，你要给哪个队列发送消息
// Send message to the exchange, so we need send the exchange name, routerKey, because, exchanges should know, what queue you want to send.
func (r RabbitmqClient) Send(ctx context.Context, exchangeName string, routerKey string, option amqp.Publishing) error {
	return r.ch.PublishWithContext(
		ctx,          // 传递的上下文
		exchangeName, // 交换机名称
		routerKey,    // 路由键，告诉交换机将消息传递给哪些队列中去
		true,
		false,
		option, // 需要发送的消息内容
	)
}

// Consume 创建消费者，获取到一个只能接收消息的通道

// <-chan amqp.Delivery 表示这个函数返回一个通道，<- 表示这个通道只能接收消息，不能给这个通道发送消息， 且通道的消息数据类型为 amqp.Delivery

// chan <- message 则表示这个函数返回一个通道，<- 表示这个通道只能发送消息，不能接收消息，且通道的消息数据类型为 message

func (r RabbitmqClient) Consume(queueName, consumeName string, autoAck bool, table amqp.Table) (<-chan amqp.Delivery, error) {

	// 不带上下文的消费者，就会一直存在，上下文消失不会影响消费者
	return r.ch.Consume(
		queueName,   // 队列名称
		consumeName, // 消费者名称
		autoAck,     // 是否自动应答, 就是收到消息后，自动告诉rabbitmq 消息已经被成功消费了，从队列中删除。 一般不需要，都是后续的业务逻辑处理成功后，手动ACK
		false,       // 是否排他，这个消费者作为 这个队列的唯一消费者，其余的消费者不可以在这个队列进行消费
		false,       // 如果为true 则消费者不会接收通过同一连接发送的消息。这对于避免消费者接收其自己发送的消息非常有用。
		false,       // 是否阻塞，如果false 消费者会立马消费消息，不等待服务器的响应
		table,
	)

	// 带上下文的 消费者，当上下文消失嘞，这个消费者也就被删除了
	//return r.ch.ConsumeWithContext(
	//	ctx,
	//	queueName,   // 队列名称
	//	consumeName, // 消费者名称
	//	autoAck,     // 是否自动应答, 就是收到消息后，自动告诉rabbitmq 消息已经被成功消费了，从队列中删除。 一般不需要，都是后续的业务逻辑处理成功后，手动ACK
	//	false,       // 是否排他，这个消费者作为 这个队列的唯一消费者，其余的消费者不可以在这个队列进行消费
	//	false,       // 如果为true 则消费者不会接收通过同一连接发送的消息。这对于避免消费者接收其自己发送的消息非常有用。
	//	false,       // 是否阻塞，如果false 消费者会立马消费消息，不等待服务器的响应
	//	table,
	//)

	//return nil, nil
}
