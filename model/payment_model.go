package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaymentID     string
	Cart          []Cart
	PaymentMethod string
}
