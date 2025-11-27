package models

import (
	"github.com/google/uuid"
)

type FileSharedTeam struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FileID uuid.UUID `gorm:"type:uuid;not null" json:"file_id"`
	TeamID uuid.UUID `gorm:"type:uuid;not null" json:"team_id"`

	File File `gorm:"foreignKey:FileID" json:"file"`
	Team Team `gorm:"foreignKey:TeamID" json:"team"`
}
