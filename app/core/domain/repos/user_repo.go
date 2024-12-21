package repos

import (
	"github.com/lai0xn/bsc-auth/db"
	"github.com/lai0xn/bsc-auth/domain/enteties"
	"github.com/lai0xn/bsc-auth/pkg/utils"
)

type userRepository struct {
	GenericRepo[enteties.User]
}

func NewUserRepository() *userRepository {
	return &userRepository{
		GenericRepo: *(new(GenericRepo[enteties.User])),
	}
}

func (r *userRepository) FindUserByEmail(email string) (enteties.User, error) {
	database := db.GetDB()
	var user enteties.User
	result := database.First(&user, "email = ?", email)
	if result.Error != nil {
		return utils.Zero[enteties.User](), result.Error
	}

	return user, nil
}

func (r *userRepository) SetPassword(id int, password string) (enteties.User, error) {
	database := db.GetDB()
	hashedPassword, err := utils.Encrypt(password)
	if err != nil {
		return utils.Zero[enteties.User](), nil
	}
	var user enteties.User
	result := database.First(&user, "id = ?", id).Update("password", hashedPassword)
	if result.Error != nil {
		return utils.Zero[enteties.User](), result.Error
	}

	return user, nil
}
