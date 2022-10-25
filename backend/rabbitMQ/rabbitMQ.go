package rabbitMQ

import (
	"fmt"

	"github.com/streadway/amqp"
)

func AddPublish(ch *amqp.Channel, q amqp.Queue) {
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
            Body:        []byte("hellow world"),	
		},
	)
	fmt.Println(err)
}

func AddQUene(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"hello", // 队列名字
		true,   // 消息是否持久化
		false,   // 不使用的时候删除队列
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	fmt.Println(err)
	return q
}

func AddExchange(ch *amqp.Channel) {
	err := ch.ExchangeDeclare(
		"tizi365_topic", // 交换机名字，需要唯一
		"topic",      // 交换机类型
		true,          // 是否持久化
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	fmt.Println(err)
}

func AddConsumer (ch *amqp.Channel, q amqp.Queue) {
	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	fmt.Println(err)

	go func() {
		for d := range msgs {
		fmt.Printf("Received a message: %s", d.Body)
		}
	}()
}

func Conn() {
	// 连接RabbitMQ Server
	//amqp://账号:密码@RabbitMQ地址:端口/

	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	fmt.Println(err)

	ch, err := conn.Channel()
	fmt.Println(err)

	AddQUene(ch)
	// AddExchange(ch)

}

