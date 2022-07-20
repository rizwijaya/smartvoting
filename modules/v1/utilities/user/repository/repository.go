package repository

import (
	"smartvoting/modules/v1/utilities/user/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindByToken(token string) (models.User, error)
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
