package service

import (
	"errors"
	"smartvoting/modules/v1/utilities/user/models"
	"smartvoting/modules/v1/utilities/user/repository"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
	NewToken(token string) (string, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) Login(input models.LoginInput) (models.User, error) {
	token := input.Token

	user, err := s.repository.FindByToken(token)
	if err != nil {
		return user, err
	}

	if user.User_id == 0 {
		return user, errors.New("Token yang ada masukkan salah!")
	}

	// err = bcrypt.CompareHashAndPassword([]byte(user.Token), []byte(token))
	// if err != nil {
	// 	return user, errors.New("Token yang ada masukkan salah!")
	// }

	return user, nil
}

func (s *service) NewToken(token string) (string, error) {
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	token = string(tokenHash)
	if err != nil {
		return token, err
	}

	return token, err
}
