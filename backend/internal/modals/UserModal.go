package modals

type Users struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"type:text;not null" json:"email"`
	Password string `gorm:"type:text;not null" json:"password"`
	Verified bool   `gorm:"type:boolean;not null" json:"verified"`
}
