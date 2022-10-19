package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID          string `gorm:"primarykey"`
	FirstName       string
	LastName        string
	Email           string
	Password        string
	ConfirmPassword string
	Token           string
	RefreshToken    string
}
