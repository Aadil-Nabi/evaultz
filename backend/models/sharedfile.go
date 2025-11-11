package models

import (
	"time"

	"github.com/google/uuid"
)

type SharedFile struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FileID     uuid.UUID `gorm:"type:uuid;not null"`
	File       File      `gorm:"foreignKey:FileID;constraint:OnDelete:CASCADE"`
	SharedWith uuid.UUID `gorm:"type:uuid;not null"`
	SharedUser User      `gorm:"foreignKey:SharedWith;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
