package usecases

import (
	"di.com/m/v2/repositories"
	"github.com/google/wire"
)

type UseCaseB struct {
	repository repositories.IUserRepository
}

func (useCaseB *UseCaseB) Get(id int) *User {
	user := useCaseB.repository.GetUserById(id)
	return mapUser(user)
}

func NewUseCaseB(repository repositories.IUserRepository) UserUseCase {
	usecase := new(UseCaseB)
	usecase.repository = repository
	return usecase
}

var HandlerSetDi = wire.NewSet(NewUseCaseB, repositories.NewUserRepositoryA)
