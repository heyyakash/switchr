package modals

import "time"

type Roles struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:text;not null;unique" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
