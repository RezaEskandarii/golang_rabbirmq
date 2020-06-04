package rabbitmq

import (
	"github.com/streadway/amqp"
	"os"
)

//
func GetQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))

	if err != nil {
		panic(err)
		//return nil, nil, nil
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
	return conn, ch, &queue
}
