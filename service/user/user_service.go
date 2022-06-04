package user

import (
	"github.com/muhangga/entity"
	web "github.com/muhangga/web/request"
)

type UserService interface {
	RegisterUser(userRequest web.RegisterRequest) (entity.User, error)
	Login(userRequest web.LoginRequest) (entity.User, error)
	IsEmailAvailable(checkEmail web.CheckEmailRequest) (bool, error)
	SaveAvatar(userID int, fileLocation string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
}


