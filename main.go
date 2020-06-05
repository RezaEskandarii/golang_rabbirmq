package main

import (
	"./queue_utils"
	"fmt"
	//"github.com/joho/godotenv"
	"strconv"
	"time"
)

func main() {

	duration, _ := time.ParseDuration(strconv.Itoa(100) + "ms")
	signal := time.Tick(duration)
	server := queue_utils.Server{}
	defer server.Conn.Close()
	defer server.Channel.Close()

	client := queue_utils.Client{}
	defer client.Conn.Close()
	defer client.Channel.Close()

	for range signal {
		server.PublishMessage("test")
		go client.Consume("test")
	}

	var a string
	fmt.Scanln(&a)
}
