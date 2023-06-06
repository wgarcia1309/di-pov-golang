package repositories

import (
	"fmt"
)

type UserRepositoryA struct {
	Configuration *Config
}

func (repository *UserRepositoryA) GetUserById(id int) *User {
	fmt.Println(repository.Configuration.Path + " from repository")
	user := new(User)
	user.Name = "Fernando Sandoval"
	user.UserName = "Fernando Sandoval"
	user.Id = id
	return user
}

func NewUserRepositoryA(config *Config) IUserRepository {
	rep := new(UserRepositoryA)
	rep.Configuration = config
	return rep
}
