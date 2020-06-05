package main

import (
	"./models"
	"bytes"
	"encoding/gob"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	_ = godotenv.Load()
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	for {
		reading := models.Message{
			Name:      "Msg1",
			Value:     "Hu",
			TimeStamp: time.Now(),
		}
		buf.Reset()
		_ = enc.Encode(&reading)

	}
}
