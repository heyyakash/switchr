package db

import (
	"errors"

	"gihtub.com/heyyakash/switchr/internal/modals"
)

func (p *PostgresStore) CreateUserProjectMap(userprojectmap *modals.UserProjectMap) error {
	err := p.DB.Create(&userprojectmap).Error
	return err
}

func (p *PostgresStore) GetUserProjectMapByUidAndPid(uid string, pid string) (modals.UserProjectMap, error) {
	var userprojectmap modals.UserProjectMap
	res := p.DB.Where("uid = ?", uid).Where("pid = ?::uuid", pid).Preload("Project").Preload("User").First(&userprojectmap)
	if res.Error != nil {
		return modals.UserProjectMap{}, res.Error
	}
	if res.RowsAffected == 0 {
		return modals.UserProjectMap{}, errors.New("no matching record found")
	}
	return userprojectmap, nil
}

func (p *PostgresStore) GetUserProjectMapByUid(uid string) ([]modals.UserProjectMap, error) {
	var userprojectmaps []modals.UserProjectMap
	res := p.DB.Where("uid = ?", uid).Preload("Project").Find(&userprojectmaps)
	if res.Error != nil {
		return []modals.UserProjectMap{}, res.Error
	}
	return userprojectmaps, nil
}