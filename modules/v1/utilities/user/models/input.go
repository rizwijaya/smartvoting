package models

type LoginInput struct {
	Token string `json:"token" form:"token" binding:"required"`
}
