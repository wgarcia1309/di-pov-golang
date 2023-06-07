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

func InitializeHandlerA() handlers.UserHandler {
	wire.Build(a.NewHandler, usecases.NewUseCaseA, repositories.NewUserRepositoryA)
	return &a.UserHandlerA{}
}

func InitializeHandlerB() handlers.UserHandler {
	wire.Build(b.NewHandler, usecases.NewUseCaseB, repositories.NewUserRepositoryB)
	return &b.UserHandlerB{}
}

func InitializeHandlerC() handlers.UserHandler {
	wire.Build(b.NewHandler, usecases.NewUseCaseA, repositories.NewUserRepositoryB)
	return &b.UserHandlerB{}
}
