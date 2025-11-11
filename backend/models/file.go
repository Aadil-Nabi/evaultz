package models

import (
	"time"

	"github.com/google/uuid"
)

type Visibility string

const (
	Private Visibility = "private"
	Shared  Visibility = "shared"
	Public  Visibility = "public"
)

type File struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null"`
	Owner       User      `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
	Name        string    `gorm:"not null"`
	S3Key       string    `gorm:"not null;uniqueIndex"`
	Size        int64
	Visibility  Visibility   `gorm:"type:varchar(20);default:'private'"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
	SharedUsers []SharedFile `gorm:"foreignKey:FileID"`
}
