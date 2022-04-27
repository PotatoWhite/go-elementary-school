package main

import (
	"fmt"
	"os"
)

func NewEvent(g Greeter) (Event, error) {
	return Event{Greeter: g}, nil
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewSpec(g Greeter, e Event, m Message) Spec {

	return Spec{
		Event:   e,
		Greeter: g,
		Message: m,
	}
}

func main() {
	spec, err := InitializeSpec("Hello Carrot")
	if err != nil {
		fmt.Printf("failed to create event: %v\n", err)
		os.Exit(2)
	}
	spec.Event.Start()
}
