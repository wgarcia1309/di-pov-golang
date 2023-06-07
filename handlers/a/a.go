package a

import (
	"di.com/m/v2/handlers"
	"di.com/m/v2/usecases"
)

type UserHandlerA struct {
	usecase usecases.UserUseCase
}

func (userHandlerA *UserHandlerA) GetData(id int) *usecases.User {
	user := userHandlerA.usecase.Get(id)
	return user
}

func NewHandler(u usecases.UserUseCaseA) handlers.UserHandlerA {
	handler := UserHandlerA{usecase: u}
	return &handler
}
