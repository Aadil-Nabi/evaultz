package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email       string    `gorm:"uniqueIndex;not null"  json:"email"`
	Password    string    `gorm:"not null"  json:"password"`
	Username    string    ` json:"username"`
	CompanyName string    ` json:"companyname"`
	CreatedAt   time.Time `gorm:"autoCreateTime"  json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"  json:"updated_at"`
	DeletedAt   time.Time ` json:"deleted_at"`
	Files       []File    `gorm:"foreignKey:OwnerID"`
}
