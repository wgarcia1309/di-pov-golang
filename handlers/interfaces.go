package handlers

import "di.com/m/v2/usecases"

type UserHandler interface {
	GetData(id int) *usecases.User
}

type UserHandlerA UserHandler
type UserHandlerB UserHandler
