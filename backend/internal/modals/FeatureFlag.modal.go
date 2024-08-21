package modals

import "time"

type Featureflag struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Flag      string    `gorm:"type:text;not null" json:"flag"`
	Fid       string    `gorm:"type:uuid;default:uuid_generate_v4();unique;not null" json:"fid"`
	Value     string    `gorm:"type:text;not null" json:"value"`
	Pid       string    `gorm:"type:uuid;not null" json:"pid"`
	CreatedBy string    `gorm:"type:uuid;not null" json:"createdBy"`
	UpdatedBy string    `gorm:"type:uuid;not null" json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Project Projects `gorm:"foreignKey:Pid;references:Pid;constraint:OnDelete:CASCADE;"`
}

type FeatureflagWithUserName struct {
	Featureflag
	FullName string `gorm:"type:text" json:"full_name"`
}
