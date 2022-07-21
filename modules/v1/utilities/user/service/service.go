package service

import (
	"errors"
	api "smartvoting/app/contracts"
	"smartvoting/modules/v1/utilities/user/models"
	"smartvoting/modules/v1/utilities/user/repository"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
	NewToken(token string) (string, error)
	GetUserAddress(id string) (string, error)
	GetUserData(address string) (models.UserData, error)
}

type service struct {
	repository repository.Repository
	blockchain *api.Api
}

func NewService(repository repository.Repository, blockchain *api.Api) *service {
	return &service{repository, blockchain}
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

func (s *service) GetUserAddress(id string) (string, error) {
	address, err := s.repository.GetUserAddress(id)
	if err != nil {
		return address, err
	}

	return address, nil
}

func (s *service) GetUserData(address string) (models.UserData, error) {
	userAddress := common.HexToAddress(address)
	userData, err := s.blockchain.DetailPemilih(&bind.CallOpts{}, userAddress)
	if err != nil {
		return userData, err
	}

	return userData, nil
}
