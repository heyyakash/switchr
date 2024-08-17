package modals

import "time"

type Featureflag struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Flag      string    `gorm:"type:text;not null" json:"flag"`
	Fid       string    `gorm:"type:uuid;default:uuid_generate_v4();unique;not null" json:"fid"`
	Value     string    `gorm:"type:text;not null" json:"value"`
	Pid       string    `gorm:"type:text;not null" json:"pid"`
	CreatedBy string    `gorm:"type:text;not null" json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
