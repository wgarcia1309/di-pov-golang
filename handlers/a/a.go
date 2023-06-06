package a

import (
	"di.com/m/v2/handlers"
	"di.com/m/v2/repositories"
	"di.com/m/v2/usecases"
	"go.uber.org/dig"
)

type UserHandlerA struct {
	usecase usecases.UserUseCase
}

func (userHandlerA *UserHandlerA) GetData(id int) *usecases.User {
	user := userHandlerA.usecase.Get(id)
	return user
}

func NewHandler() handlers.UserHandler {
	container := dig.New()
	container.Provide(repositories.NewUserRepositoryA)
	container.Provide(usecases.NewUseCaseA)
	var usecase usecases.UserUseCase
	container.Invoke(func(u usecases.UserUseCase) {
		usecase = u
	})

	handler := UserHandlerA{usecase: usecase}
	return &handler
}
