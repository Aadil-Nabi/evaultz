package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID       uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OwnerID  uuid.UUID  `gorm:"type:uuid;not null" json:"owner_id"`
	TenantID uuid.UUID  `gorm:"type:uuid;not null" json:"tenant_id"`
	TeamID   *uuid.UUID `gorm:"type:uuid" json:"team_id"`

	FileName   string `gorm:"not null" json:"file_name"`
	FileKey    string `gorm:"not null;uniqueIndex" json:"file_key"`
	MimeType   string `json:"mime_type"`
	SizeBytes  int64  `json:"size_bytes"`
	Visibility string `gorm:"type:text;not null;check:visibility IN ('public','private','shared')" json:"visibility"`

	Bucket string `gorm:"not null" json:"bucket"`
	Region string `gorm:"not null" json:"region"`

	UploadedAt time.Time `json:"uploaded_at"`

	Owner  User   `gorm:"foreignKey:OwnerID" json:"owner"`
	Tenant Tenant `gorm:"foreignKey:TenantID" json:"tenant"`
	Team   *Team  `gorm:"foreignKey:TeamID" json:"team"`

	SharedUsers   []FileSharedUser   `json:"shared_users"`
	SharedTeams   []FileSharedTeam   `json:"shared_teams"`
	SharedTenants []FileSharedTenant `json:"shared_tenants"`
}
