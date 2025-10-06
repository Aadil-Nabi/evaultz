package models

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	NameOnCard string
	CardNumber string
	Cvv        string
	Expiry     string
	UserID     uint  // foreign key column
	User       *User `gorm:"foreignKey:UserID"`
}
