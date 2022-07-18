package service

import (
	"smartvoting/modules/v1/utilities/user/repository"
)

type Service interface {
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
