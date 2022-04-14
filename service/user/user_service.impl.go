package service

import (
	"github.com/muhangga/entity"
	"github.com/muhangga/helper"
	"golang.org/x/crypto/bcrypt"

	model "github.com/muhangga/model/request"
	repository "github.com/muhangga/repository/user"
)

type userService struct {
	userRepository repository.Repository
}

func NewService(userRepository repository.Repository) *userService {
	return &userService{userRepository}
}

func (s *userService) RegisterUser(userRequest model.RegisterRequest) (entity.User, error) {
	user := entity.User{}
	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Occupation = userRequest.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = helper.HashPassword(passwordHash)
	user.Role = "user"

	users, err := s.userRepository.Save(user)
	if err != nil {
		return users, err
	}

	return users, nil
}
