package b

import (
	"di.com/m/v2/handlers"
	"di.com/m/v2/repositories"
	"di.com/m/v2/usecases"
	"go.uber.org/dig"
)

type UserHandlerB struct {
	usecase usecases.UserUseCase
}

func (userHandlerB *UserHandlerB) GetData(id int) *usecases.User {
	user := userHandlerB.usecase.Get(id)
	return user
}

func NewHandler() handlers.UserHandler {
	container := dig.New()

	container.Provide(repositories.NewUserRepositoryB)
	container.Provide(usecases.NewUseCaseA)
	var usecase usecases.UserUseCase
	container.Invoke(func(u usecases.UserUseCase) {
		usecase = u
	})

	handler := UserHandlerB{usecase: usecase}
	return &handler
}
