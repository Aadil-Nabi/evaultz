package models

import "gorm.io/gorm"

type Files struct {
	gorm.Model
	Key      string
	Location string
}
