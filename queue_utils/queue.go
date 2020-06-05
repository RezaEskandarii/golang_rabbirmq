package queue_utils

import (
	"../utils"
	"github.com/streadway/amqp"
)

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	utils.FailOnError(err, "Failed in connection")
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to get channel")
	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(name,
		false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare queue")
	return &q
}
