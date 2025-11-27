package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OwnerID  uuid.UUID  `gorm:"type:uuid;not null"`
	TenantID uuid.UUID  `gorm:"type:uuid;not null"`
	TeamID   *uuid.UUID `gorm:"type:uuid"`

	FileName   string `gorm:"not null"`
	StorageKey string `gorm:"not null"`
	Size       int64
	URL        string
	MimeType   string
	Visibility string `gorm:"type:text;default:'private'"`

	CreatedAt time.Time
	UpdatedAt time.Time

	Owner  User   `gorm:"foreignKey:OwnerID"`
	Tenant Tenant `gorm:"foreignKey:TenantID"`
	Team   *Team  `gorm:"foreignKey:TeamID"`
}
