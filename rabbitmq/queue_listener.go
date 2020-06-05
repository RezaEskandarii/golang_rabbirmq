package rabbitmq

import "github.com/streadway/amqp"

type QueueListener struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}
