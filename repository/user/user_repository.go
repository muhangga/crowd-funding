package repository

import (
	"github.com/muhangga/entity"
)

type Repository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}