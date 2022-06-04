package user

import (
	"github.com/muhangga/entity"
	"github.com/muhangga/model/request"
)

type UserService interface {
	RegisterUser(userRequest model.RegisterRequest) (entity.User, error)
	Login(userRequest model.LoginRequest) (entity.User, error)
	IsEmailAvailable(checkEmail model.CheckEmailRequest) (bool, error)
	SaveAvatar(userID int, fileLocation string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
}


