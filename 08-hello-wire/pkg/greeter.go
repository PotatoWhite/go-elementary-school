package main

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}
