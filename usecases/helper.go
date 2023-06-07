package usecases

import (
	"di.com/m/v2/repositories"
)

func mapUser(model *repositories.User) *User {
	newUser := new(User)
	newUser.Name = model.Name
	newUser.UserName = model.UserName
	newUser.Id = model.Id
	return newUser
}
