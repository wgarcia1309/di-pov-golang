package repositories

type IUserRepository interface {
	GetUserById(id int) *User
}
type IUserRepositoryA IUserRepository
type IUserRepositoryB IUserRepository
