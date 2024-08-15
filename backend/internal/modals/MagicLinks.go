package modals

import "time"

type MagicLink struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt" `
	ExpiresAt time.Time `gorm:"not null" json:"expiresAt"`
}
