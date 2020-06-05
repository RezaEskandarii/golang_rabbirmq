package queue_utils

import (
	"../database"
	"../models"
	"../utils"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

type Client struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (c *Client) Consume(queueName string) {

	db, err := database.GetConnection()
	if err != nil {
		return
	}
	defer db.Close()
	fmt.Println(db)
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_HOST"))
	utils.FailOnError(err, "")
	c.Conn = conn

	c.Channel, err = c.Conn.Channel()
	utils.FailOnError(err, "")

	msgs, err := c.Channel.Consume(queueName, "", true, false, false, false, nil)
	utils.FailOnError(err, "")
	var a string
	for msg := range msgs {
		//fmt.Println(string(msg.Body))
		buf := bytes.NewBuffer([]byte(msg.Body))
		decode := gob.NewDecoder(buf)
		message := models.Message{}
		if err := decode.Decode(&message); err != nil {
			utils.FailOnError(err, err.Error())
		}
		db.Create(&message)
	}
	fmt.Scanln(&a)
}
