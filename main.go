package main

import (
	"fmt"

	"di.com/m/v2/handlers"
	"di.com/m/v2/handlers/a"
	"di.com/m/v2/handlers/b"
	"di.com/m/v2/repositories"
	"di.com/m/v2/usecases"
	"go.uber.org/dig"
)

func main() {

	container := dig.New()
	container.Provide(repositories.NewUserRepositoryA)
	container.Provide(usecases.NewUseCaseA)
	container.Provide(a.NewHandler)

	container.Provide(repositories.NewUserRepositoryB)
	container.Provide(usecases.NewUseCaseB)
	container.Provide(b.NewHandler)
	err := container.Invoke(func(h handlers.UserHandlerA) {
		user := h.GetData(123)
		fmt.Println(user)
	})
	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(h handlers.UserHandlerB) {
		user := h.GetData(456)
		fmt.Println(user)
	})
	if err != nil {
		panic(err)
	}

}
