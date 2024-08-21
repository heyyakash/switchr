package modals

import (
	"time"

	"gihtub.com/heyyakash/switchr/internal/constants"
)

type UserProjectMap struct {
	Uid       string         `gorm:"type:uuid;not null;primaryKey" json:"uid"`
	Pid       string         `gorm:"type:uuid;not null;primaryKey" json:"pid"`
	Role      constants.Role `gorm:"not null" json:"role_id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`

	// Foreign key relationships
	Project Projects `gorm:"foreignKey:Pid;references:Pid;constraint:OnDelete:CASCADE;"`
	User    Users    `gorm:"foreignKey:Uid;references:Uid;constraint:OnDelete:CASCADE;"`
}
