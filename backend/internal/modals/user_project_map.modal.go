package modals

import (
	"time"

	"gihtub.com/heyyakash/switchr/internal/constants"
)

type UserProjectMap struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	Uid       string         `gorm:"type:text;not null" json:"uid"`
	Pid       string         `gorm:"not null;constraint:OnDelete:CASCADE;" json:"pid"`
	Role      constants.Role `gorm:"not null" json:"role_id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	// Define the foreign key relationship
	Project Projects `gorm:"foreignKey:Pid;references:Id"`
	User    Users    `gorm:"foreignKey:Uid;references:Uid"`
}
