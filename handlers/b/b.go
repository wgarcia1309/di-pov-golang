package b

import (
	"di.com/m/v2/handlers"
	"di.com/m/v2/usecases"
)

type UserHandlerB struct {
	usecase usecases.UserUseCase
}

func (userHandlerB *UserHandlerB) GetData(id int) *usecases.User {
	user := userHandlerB.usecase.Get(id)
	return user
}

func NewHandler(u usecases.UserUseCaseB) handlers.UserHandlerB {
	handler := UserHandlerB{usecase: u}
	return &handler
}
