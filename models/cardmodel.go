package models

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Name       string
	CardNumber string
	Cvv        string
	Expiry     string
}
