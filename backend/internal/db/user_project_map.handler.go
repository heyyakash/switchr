package db

import "gihtub.com/heyyakash/switchr/internal/modals"

func (p *PostgresStore) CreateUserProjectMap(userprojectmap *modals.UserProjectMap) error {
	err := p.DB.Create(&userprojectmap).Error
	return err
}

func (p *PostgresStore) GetUserProjectMapByUidAndPid(uid string, pid string) (modals.UserProjectMap, error) {
	var userprojectmap modals.UserProjectMap
	res := p.DB.Where("uid = ?", uid).Where("pid = ?", pid).First(&userprojectmap)
	if res.Error != nil {
		return modals.UserProjectMap{}, res.Error
	}
	return userprojectmap, nil
}
