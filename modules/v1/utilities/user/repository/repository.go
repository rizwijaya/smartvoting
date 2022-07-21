package repository

import (
	"smartvoting/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByToken(token string) (models.User, error)
	GetUserAddress(id string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByToken(token string) (models.User, error) {
	var user models.User
	err := r.db.Where("token = ?", token).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) GetUserAddress(id string) (string, error) {
	var address string
	err := r.db.Raw("SELECT address FROM users WHERE user_id = ?", id).Find(&address).Error
	if err != nil {
		return address, err
	}
	return address, nil
}
