package utils

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

func FailOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s : %s", message, err)
		panic(fmt.Sprintf("%s: %s", message, err))
	}
}
