package repositories

type IUserRepository interface {
	GetUserById(id int) *User
}
