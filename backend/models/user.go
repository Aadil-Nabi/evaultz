package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TenantID  uuid.UUID  `gorm:"type:uuid;not null" json:"tenant_id"`
	TeamID    *uuid.UUID `gorm:"type:uuid" json:"team_id"` // optional (user may not belong to a team)
	Email     string     `gorm:"uniqueIndex;not null" json:"email"`
	Username  string     `gorm:"not null" json:"username"`
	Password  string     `gorm:"not null" json:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
	Team   *Team  `gorm:"foreignKey:TeamID" json:"team"`

	Files []File `gorm:"foreignKey:OwnerID" json:"files"`
}
