package usecases

import (
	"di.com/m/v2/repositories"
)

type UseCaseA struct {
	repository repositories.IUserRepository
}

func (useCaseA *UseCaseA) Get(id int) *User {
	user := useCaseA.repository.GetUserById(id)
	return mapUser(user)
}

func NewUseCaseA(repository repositories.IUserRepositoryA) UserUseCaseA {
	usecase := new(UseCaseA)
	usecase.repository = repository
	return usecase
}
