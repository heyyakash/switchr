package modals

import (
	"time"

	"gihtub.com/heyyakash/switchr/internal/constants"
)

type UserProjectMap struct {
	Id        uint           `gorm:"primaryKey" json:"id"`
	Uid       string         `gorm:"type:text;not null" json:"uid"`
	ProjectId uint           `gorm:"not null;constraint:OnDelete:CASCADE;" json:"project_id"`
	Role      constants.Role `gorm:"not null" json:"role_id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	// Define the foreign key relationship
	Project Projects `gorm:"foreignKey:ProjectId;references:Id"`
	User    Users    `gorm:"foreignKey:Uid;references:Uid"`
}
