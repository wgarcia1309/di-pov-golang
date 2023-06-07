//go:build wireinject
// +build wireinject

package main

import (
	"di.com/m/v2/handlers"
	"di.com/m/v2/handlers/a"
	"di.com/m/v2/handlers/b"
	"di.com/m/v2/repositories"
	"di.com/m/v2/usecases"
	"github.com/google/wire"
)

var handlerSetA = wire.NewSet(a.NewHandler, usecases.NewUseCaseA, repositories.NewUserRepositoryA)
var handlerSetB = wire.NewSet(a.NewHandler, usecases.NewUseCaseA, repositories.NewUserRepositoryA)
var handlerSetC = wire.NewSet(a.NewHandler, usecases.NewUseCaseA, repositories.NewUserRepositoryA)

func InitializeHandlerA() handlers.UserHandler {
	wire.Build(handlerSetA)
	return &a.UserHandlerA{}
}

func InitializeHandlerB() handlers.UserHandler {
	wire.Build(handlerSetB)
	return &b.UserHandlerB{}
}

func InitializeHandlerC() handlers.UserHandler {
	wire.Build(handlerSetC)
	return &b.UserHandlerB{}
}
