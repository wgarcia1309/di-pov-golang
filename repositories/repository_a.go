package repositories

type UserRepositoryA struct {
}

func (repository *UserRepositoryA) GetUserById(id int) *User {
	user := new(User)
	user.Name = "Fernando Sandoval"
	user.UserName = "FS"
	user.Id = id
	return user
}

func NewUserRepositoryA() IUserRepository {
	rep := new(UserRepositoryA)
	return rep
}
