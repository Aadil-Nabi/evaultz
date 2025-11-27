package models

import (
	"github.com/google/uuid"
)

type FileSharedUser struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FileID uuid.UUID `gorm:"type:uuid;not null" json:"file_id"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	File File `gorm:"foreignKey:FileID" json:"file"`
	User User `gorm:"foreignKey:UserID" json:"user"`
}
