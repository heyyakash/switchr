package modals

import "time"

type Projects struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null" json:"name"`
	CreatedBy string    `gorm:"type:text;not null" json:"createdBy"`
	Pid       string    `gorm:"type:uuid;default:uuid_generate_v4();unique;not null" json:"pid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
