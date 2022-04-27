package main

import (
	"fmt"
)

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
