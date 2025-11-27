package models

import (
	"github.com/google/uuid"
)

type FileSharedTenant struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FileID   uuid.UUID `gorm:"type:uuid;not null" json:"file_id"`
	TenantID uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`

	File   File   `gorm:"foreignKey:FileID" json:"file"`
	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
}
