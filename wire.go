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
var handlerSetB = wire.NewSet(b.NewHandler, usecases.NewUseCaseB, repositories.NewUserRepositoryB)
var handlerSetC = wire.NewSet(b.NewHandler, usecases.NewUseCaseA, repositories.NewUserRepositoryB)

func InitializeHandlerA() handlers.UserHandler {
	wire.Build(handlerSetA)
	return nil
}

func InitializeHandlerB() handlers.UserHandler {
	wire.Build(handlerSetB)
	return nil
}

func InitializeHandlerC() handlers.UserHandler {
	wire.Build(handlerSetC)
	return nil
}

func InitializeHandlerD() handlers.UserHandler {
	wire.Build(b.NewHandler, usecases.HandlerSetDi)
	return nil
}
