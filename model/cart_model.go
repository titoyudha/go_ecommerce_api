package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartID     string
	User       []User
	Items      []Items
	TotalPrice int
	Status     string
}
