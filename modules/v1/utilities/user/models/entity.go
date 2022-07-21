package models

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type User struct {
	User_id      int
	Role_id      int
	Token        string
	Address      string
	Date_updated time.Time
	Date_created time.Time
}

type UserData struct {
	VoterAddress common.Address
	Name         string
	HomeAddress  string
	Phone        string
	IsVerified   bool
	HasVoted     bool
	IsRegistered bool
}
