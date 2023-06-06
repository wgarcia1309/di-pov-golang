package repositories

import (
	"fmt"
)

type UserRepositoryB struct {
	Configuration *Config
}

func (repository *UserRepositoryB) GetUserById(id int) *User {
	fmt.Println(repository.Configuration.Path + " from repository")
	user := new(User)
	user.Name = "Juan Valdez"
	user.UserName = "Juan Valdez"
	user.Id = id
	return user
}

func NewUserRepositoryB(config *Config) IUserRepository {
	rep := new(UserRepositoryB)
	rep.Configuration = config
	return rep
}
