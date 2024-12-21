package user

import (
	"github.com/lai0xn/bsc-auth/app/user/dto"
	"github.com/lai0xn/bsc-auth/domain/enteties"
	"github.com/lai0xn/bsc-auth/pkg/rabittmq"
	"github.com/lai0xn/bsc-auth/pkg/utils"
)

type UserService interface {
	Authenticate(email, password string) (enteties.User, error)
	Register(dto dto.UserDto) error
}

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Authenticate(email string, password string) (enteties.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return utils.Zero[enteties.User](), err
	}

	err = utils.CheckPassword(user.Password, password)

	if err != nil {
		return utils.Zero[enteties.User](), err
	}

	return user, nil
}

func (s *userService) Register(dto dto.UserDto) error {
	hashedPassword, err := utils.Encrypt(dto.Password)
	if err != nil {
		return err
	}
	_, err = s.repo.Create(enteties.User{
		Email:        dto.Email,
		FirstName:    dto.FirstName,
		LastName:     dto.FirstName,
		Password:     hashedPassword,
		IsStaff:      false,
		IsActive:     true,
		DateOfBirth:  dto.DateOfBirth,
		PlaceOfBirth: dto.PlaceOfBirth,
		Gender:       dto.Gender,
	})
	if err != nil {
		return err
	}
	q := rabittmq.NewQueue("users")
	err = q.Publish(dto)
	if err != nil {
		return err
	}
	return nil
}
