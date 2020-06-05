package models

import (
	"encoding/gob"
	"strings"
	"time"
)

type Message struct {
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	TimeStamp time.Time `json:"time_stamp"`
}

// check message validation
func (m *Message) Valid() bool {
	if strings.TrimSpace(m.Name) == "" && strings.TrimSpace(m.Value) == "" {
		return false
	}
	return true
}

// Register records a type, identified by a value for that type, under its
// internal type name. That name will identify the concrete type of a value
// sent or received as an interface variable. Only types that will be
// transferred as implementations of interface values need to be registered.
// Expecting to be used only during initialization, it panics if the mapping
// between types and names is not a bijection.
func init() {
	gob.Register(&Message{})
}
