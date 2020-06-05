package queue_utils

import (
	"../models"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

type QueueListener struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	sources    map[string]<-chan amqp.Delivery
}

// return new queue listener
func NewQueueListener() *QueueListener {
	qListener := QueueListener{
		sources: make(map[string]<-chan amqp.Delivery),
	}
	qListener.connection, qListener.channel = GetChannel(os.Getenv("RABBITMQ_HOST"))
	return &qListener
}

func (ql *QueueListener) ListenForNewSource() {
	queue := GetQueue("Reza", ql.channel)
	ql.channel.QueueBind(queue.Name, "", "amq.fanout", false, nil)

	messages, _ := ql.channel.Consume(queue.Name, "", false, false,
		false, false, nil)

	for msg := range messages {
		fmt.Println(msg.Body)
		sourceChan, _ := ql.channel.Consume(string(msg.Body), "",
			true, false, false, false, nil)
		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(deliveries <-chan amqp.Delivery) {
	for msg := range deliveries {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		messageModel := models.Message{}
		d.Decode(&messageModel)
		fmt.Printf("reaceived message: %s \n", string(msg.Body))
	}
}
