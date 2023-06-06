package handlers

import (
	"di.com/m/v2/usecases"
)

type UserHandlerA struct {
	usecase usecases.UserUseCase
}

func (userHandlerA *UserHandlerA) GetData(id int) *usecases.User {
	user := userHandlerA.usecase.Get(id)
	return user
}

func NewUserHandlerA(userUsecae usecases.UserUseCase) UserHandler {
	handler := UserHandlerA{usecase: userUsecae}
	return &handler
}
