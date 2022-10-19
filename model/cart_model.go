package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartID     string `gorm:"primarykey"`
	User       []User
	Items      []Items
	TotalPrice int
	Status     string
}
