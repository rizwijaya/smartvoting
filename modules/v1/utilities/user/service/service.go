package service

import (
	"errors"
	"log"
	api "smartvoting/app/contracts"
	"smartvoting/modules/v1/utilities/user/models"
	"smartvoting/modules/v1/utilities/user/repository"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/labstack/gommon/random"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
	NewToken(token string) (string, error)
	GetUserAddress(id string) (string, error)
	GetUserData(address string) (models.UserData, error)
	NewUser(input models.NewUser) (*types.Transaction, error)
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

func (s *service) NewUser(input models.NewUser) (*types.Transaction, error) {
	user := models.User{}
	user.Role_id = 2
	user.Token = random.String(14)
	user.Address = input.Address

	err := s.repository.NewUser(user)
	if err != nil {
		return nil, err
	}

	userAddress := common.HexToAddress(input.Address)
	userData, err := s.blockchain.PostAdminPemilihBaru(&bind.TransactOpts{}, userAddress, input.Nama, input.Alamat, input.Nomor)
	if err != nil {
		log.Println(err) //Print log error
		return userData, err
	}
	return userData, nil
}
