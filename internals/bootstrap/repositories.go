package bootstrap

import (
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/db"
	"github.com/iacopoghilardi/golang-backend-boilerplate/internals/repositories"
)

type Repositories struct {
	UserRepository *repositories.UserRepository
}

func SetupRepositories() *Repositories {
	return &Repositories{
		UserRepository: repositories.NewUserRepository(db.GetDB()),
	}
}
