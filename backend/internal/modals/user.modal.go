package modals

import "time"

type Users struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Uid       string    `gorm:"type:uuid;default:uuid_generate_v4();unique;not null" json:"uid"`
	Email     string    `gorm:"type:text;not null" json:"email"`
	FullName  string    `gorm:"type:text;not null" json:"fullname"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	Verified  bool      `gorm:"type:boolean;not null" json:"verified"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
