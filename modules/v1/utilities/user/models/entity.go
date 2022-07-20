package models

import "time"

type User struct {
	User_id      int
	Role_id      int
	Token        string
	Address      string
	Date_updated time.Time
	Date_created time.Time
}
