package rabbitmq

import (
	"github.com/streadway/amqp"
	"os"
)

//
func GetQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))

	if err != nil {
		panic(err)
		return nil, nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	queue, err := ch.QueueDeclare("Hello",
		false, false,
		false, false, nil)
	if err != nil {
		panic(err)
	}
	return conn, ch, &queue, err
}
