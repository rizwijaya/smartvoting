package models

type NewElection struct {
	Server     string `json:"server" form:"server" binding:"required"`
	Organisasi string `json:"organisasi" form:"organisasi" binding:"required"`
	Email      string `json:"email" form:"email" binding:"required"`
	Nama       string `json:"nama" form:"nama" binding:"required"`
	Jabatan    string `json:"jabatan" form:"jabatan" binding:"required"`
	Nama_acara string `json:"nama_acara" form:"nama_acara" binding:"required"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi" binding:"required"`
}
