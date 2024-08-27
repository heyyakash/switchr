package db

import (
	"gihtub.com/heyyakash/switchr/internal/modals"
	"gorm.io/gorm"
)

func (p *PostgresStore) CreateAccount(user *modals.Users) error {
	user.Verified = false
	if res := p.DB.Create(&user); res.Error != nil {
		return res.Error
	}

	return nil
}

func (p *PostgresStore) GetUserById(id uint) (modals.Users, error) {
	var user modals.Users
	if result := p.DB.Select("*").Omit("password").Where("id = ?", id).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil

}
func (p *PostgresStore) GetUserByUid(uid string) (modals.Users, error) {
	var user modals.Users
	if result := p.DB.Omit("password").Where("uid = ?", uid).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
func (p *PostgresStore) GetUserByUidWithPassword(uid string) (modals.Users, error) {
	var user modals.Users
	if result := p.DB.Where("uid = ?", uid).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (p *PostgresStore) GetUserByEmail(email string) (modals.Users, error) {
	var user modals.Users
	if result := p.DB.Select("*").Where("email = ?", email).First(&user); result.Error != nil || result.RowsAffected == 0 {
		return user, result.Error
	}

	return user, nil

}

func (p *PostgresStore) UpdateUser(user *modals.Users, uid string) error {
	return p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(modals.Users{}).Where("uid = ?", uid).Updates(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func (p *PostgresStore) EmailExists(email string) bool {
	var user modals.Users
	res := p.DB.Where("email = ?", email).First(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}
