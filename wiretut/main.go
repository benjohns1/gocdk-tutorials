package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Message string
type Message string

// Greeter struct
type Greeter struct {
	Message
	Grumpy bool
}

// Event struct
type Event struct {
	Greeter
}

// NewMessage creates a new message string
func NewMessage(phrase string) Message {
	return Message(phrase)
}

// NewGreeter creates a new message struct
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Greet returns a message
func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// NewEvent creates a new event struct
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

// Start greets via std out
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent("Hello!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
