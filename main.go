package main

import (
	"flag"
	"github.com/joho/godotenv"
)

var (
	queue = flag.String("queue", "hello-world", "queue name")
)

func main() {
	godotenv.Load()
	flag.Parse()
}
