package usecases

type UserUseCase interface {
	Get(id int) *User
}
