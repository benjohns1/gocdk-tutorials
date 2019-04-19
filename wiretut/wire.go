//+build wireinject

package main

import "github.com/google/wire"

// Initialize event given a phrase
func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
