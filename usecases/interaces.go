package usecases

type UserUseCase interface {
	Get(id int) *User
}

type UserUseCaseA UserUseCase
type UserUseCaseB UserUseCase
