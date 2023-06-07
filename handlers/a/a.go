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

func NewHandler(u usecases.UserUseCase) handlers.UserHandler {
	handler := UserHandlerA{usecase: u}
	return &handler
}
