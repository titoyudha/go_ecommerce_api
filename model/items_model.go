package model

import "gorm.io/gorm"

type Items struct {
	gorm.Model
	ItemID      string
	Title       string
	Description string
	ImageURL    string
	Price       int
}
