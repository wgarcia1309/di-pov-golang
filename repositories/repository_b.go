package repositories

type UserRepositoryB struct {
}

func (repository *UserRepositoryB) GetUserById(id int) *User {
	user := new(User)
	user.Name = "Juan Valdez"
	user.UserName = "JV"
	user.Id = id
	return user
}

func NewUserRepositoryB() IUserRepositoryB {
	rep := new(UserRepositoryB)
	return rep
}
