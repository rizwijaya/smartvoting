package repository

import (
	"smartvoting/modules/v1/utilities/voting/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetServer(status int, event string) ([]models.Server, error)
	GetServerKey(id string) (string, error)
	UpdateServer(input models.NewElection, status int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetServer(status int, event string) ([]models.Server, error) {
	server := []models.Server{}
	err := r.db.Raw("SELECT server_id, server_name, status, event FROM server WHERE status = ?"+event+";", status).Find(&server).Error
	if err != nil {
		return server, err
	}
	return server, nil
}

func (r *repository) GetServerKey(id string) (string, error) {
	var key string
	err := r.db.Raw("SELECT secretkey FROM server WHERE server_id = " + id).Find(&key).Error
	if err != nil {
		return "", err
	}
	return key, nil
}

func (r *repository) UpdateServer(input models.NewElection, status int) error {
	err := r.db.Exec("UPDATE server SET status = ?, event = ? WHERE server_id = ?;", status, input.Deskripsi, input.Server).Error
	if err != nil {
		return err
	}
	return nil
}
