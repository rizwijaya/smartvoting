package models

type LoginInput struct {
	Token string `json:"token" form:"token" binding:"required"`
}

type NewUser struct {
	Address string `json:"address" form:"address" binding:"required"`
	Nama    string `json:"nama" form:"nama" binding:"required"`
	Alamat  string `json:"alamat" form:"alamat" binding:"required"`
	Nomor   string `json:"nomor" form:"nomor" binding:"required"`
}
