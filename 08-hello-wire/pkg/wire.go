//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeSpec(phrase string) (Spec, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage, NewSpec)
	return Spec{}, nil
}
