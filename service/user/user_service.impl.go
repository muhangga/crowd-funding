package service

import (
	"errors"

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

	user := entity.User{
		Name:    userRequest.Name,
		Email:   userRequest.Email,
		Occupation: userRequest.Occupation,
		PasswordHash: helper.HashPassword([]byte(userRequest.Password)),
		Role: "user",
	}

	users, err := s.userRepository.Save(user)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) Login(loginRequest model.LoginRequest) (entity.User, error) {

	email := loginRequest.Email
	password := loginRequest.Password

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) 
	if err != nil {
		return user, errors.New("password not match")
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(checkEmailRequest model.CheckEmailRequest) (bool, error) {

	email := checkEmailRequest.Email

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}