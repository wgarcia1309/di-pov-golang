package main

import (
	"fmt"

	"di.com/m/v2/handlers"
	"di.com/m/v2/repositories"
	"di.com/m/v2/usecases"

	"go.uber.org/dig"
)

func NewConfig() *repositories.Config {
	return &repositories.Config{
		Path: "./example.db",
	}
}

func main() {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(repositories.NewUserRepositoryA)
	container.Provide(usecases.NewUseCaseA)
	container.Provide(handlers.NewUserHandlerA)

	err := container.Invoke(func(handler handlers.UserHandler) {
		user := handler.GetData(987)
		fmt.Println(user)
	})

	if err != nil {
		panic(err)
	}
	container2 := dig.New()

	container2.Provide(NewConfig)
	container2.Provide(repositories.NewUserRepositoryB)
	container2.Provide(usecases.NewUseCaseA)
	container2.Provide(handlers.NewUserHandlerA)
	err = container2.Invoke(func(handler handlers.UserHandler) {
		user := handler.GetData(458)
		fmt.Println(user)
	})

	if err != nil {
		panic(err)
	}

}
