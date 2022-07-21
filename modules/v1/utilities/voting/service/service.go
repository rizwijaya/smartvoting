package service

import (
	"log"
	api "smartvoting/app/contracts"
	"smartvoting/modules/v1/utilities/voting/models"
	"smartvoting/modules/v1/utilities/voting/repository"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Service interface {
	GetServerKey(id string) (string, error)
	CreateElection(input models.NewElection, key *bind.TransactOpts) (*types.Transaction, error)
	GetServer(status int, event string) ([]models.Server, error)
}

type service struct {
	repository repository.Repository
	blockchain *api.Api
}

func NewService(repository repository.Repository, blockchain *api.Api) *service {
	return &service{repository, blockchain}
}

func (s *service) GetServerKey(id string) (string, error) {
	server, err := s.repository.GetServerKey(id)
	if err != nil {
		log.Println(err)
		return server, err
	}
	return server, nil
}

func (s *service) CreateElection(input models.NewElection, key *bind.TransactOpts) (*types.Transaction, error) {
	//Buat Event Pemilihan di jaringan blockchain
	election, err := s.blockchain.PostAdminCreatePemilihan(key, input.Nama, input.Email, input.Jabatan, input.Deskripsi, input.Nama_acara, input.Organisasi)
	if err != nil {
		log.Println(err)
		return election, err
	}
	//Update server aktif
	err = s.repository.UpdateServer(input, 1)
	if err != nil {
		log.Println(err)
		return election, err
	}

	return election, nil
}

func (s *service) GetServer(status int, event string) ([]models.Server, error) {
	server, err := s.repository.GetServer(status, event)
	if err != nil {
		return server, err
	}
	return server, nil
}
