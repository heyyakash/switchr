package db

import "gihtub.com/heyyakash/switchr/internal/modals"

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

func (p *PostgresStore) GetUserByEmail(email string) (modals.Users, error) {
	var user modals.Users
	if result := p.DB.Select("*").Where("email = ?", email).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil

}

func (p *PostgresStore) UpdateUser(user *modals.Users) error {
	result := p.DB.Save(&user)
	return result.Error
}
