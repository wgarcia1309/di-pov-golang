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

func mapUser(model *repositories.User) *User {
	newUser := new(User)
	newUser.Name = model.Name
	newUser.UserName = model.UserName
	newUser.Id = model.Id
	return newUser
}

func NewUseCaseA(repository repositories.IUserRepository) UserUseCase {
	usecase := new(UseCaseA)
	usecase.repository = repository
	return usecase
}
