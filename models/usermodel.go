package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
}

// gorm.Model
// FirstName   string `json:"firstname" validate:"required"`
// LastName    string `json:"lastname"`
// Email       string `json:"email" validate:"email, required" gorm:"unique"`
// Password    string `json:"password" validate:"required"`
// Phone       string `json:"phone" validate:"required"`
// Token       string `json:"token"`
// RefreshType string `json:"refreshtoken"`
// UserType    string `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
// UserId      string `json:"user_id"`
