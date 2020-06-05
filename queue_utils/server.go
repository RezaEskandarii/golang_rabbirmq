package queue_utils

import (
	"../models"
	"../utils"
	"bytes"
	"encoding/gob"
	"github.com/streadway/amqp"
	"os"
	"syreclabs.com/go/faker"
)

type Server struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   *amqp.Queue
}

func (s *Server) PublishMessage(queueName string) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))
	utils.FailOnError(err, "")

	s.Conn = conn

	s.Channel, err = s.Conn.Channel()
	utils.FailOnError(err, "")

	b := new(bytes.Buffer)
	encoder := gob.NewEncoder(b)
	message := models.Message{
		Name:      faker.Name().Name(),
		Value:     faker.RandomString(25),
		TimeStamp: faker.Time().Forward(5),
	}
	b.Reset()
	encoder.Encode(message)
	msg := amqp.Publishing{
		Body: b.Bytes(),
	}

	q, err := s.Channel.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	utils.FailOnError(err, "")

	s.Queue = &q

	er := s.Channel.Publish("", s.Queue.Name, false, false, msg)
	utils.FailOnError(er, "")

}
