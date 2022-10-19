package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentID     string `gorm:"primarykey"`
	Cart          []Cart
	PaymentMethod string
}
