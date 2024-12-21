package user

import (
	"github.com/lai0xn/bsc-auth/domain/enteties"
	"github.com/lai0xn/bsc-auth/domain/repos"
)

type UserRepo interface {
	FindUserByEmail(email string) (enteties.User, error)
	SetPassword(id int, password string) (enteties.User, error)
	repos.GenericRepoI[enteties.User]
}
