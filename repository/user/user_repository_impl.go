package user

import (
	"github.com/muhangga/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user entity.User) (entity.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}