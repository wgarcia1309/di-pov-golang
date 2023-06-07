package usecases

import (
	"di.com/m/v2/repositories"
)

type UseCaseB struct {
	repository repositories.IUserRepository
}

func (useCaseB *UseCaseB) Get(id int) *User {
	user := useCaseB.repository.GetUserById(id)
	return mapUser(user)
}

func NewUseCaseB(repository repositories.IUserRepositoryB) UserUseCaseB {
	usecase := new(UseCaseB)
	usecase.repository = repository
	return usecase
}
