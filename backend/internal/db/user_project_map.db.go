package db

import (
	"errors"

	"gihtub.com/heyyakash/switchr/internal/constants"
	"gihtub.com/heyyakash/switchr/internal/modals"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (p *PostgresStore) GetMembersByPid(pid string) ([]modals.UserProjectMap, error) {
	var userProjectMaps []modals.UserProjectMap
	res := p.DB.Where("pid = ?", pid).Preload("User").Find(&userProjectMaps)
	if res.Error != nil {
		return []modals.UserProjectMap{}, res.Error
	}
	if res.RowsAffected == 0 {
		return []modals.UserProjectMap{}, errors.New("not found")
	}
	return userProjectMaps, nil
}

func (p *PostgresStore) DeleteUserProjectMapByUidPid(uid string, pid string) error {
	return p.DB.Transaction(func(tx *gorm.DB) error {
		var record modals.UserProjectMap
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&record, "uid = ? AND pid = ?", uid, pid).Error; err != nil {
			return err
		}
		if err := tx.Delete(&record).Error; err != nil {
			return err
		}

		return nil
	})
}

func (p *PostgresStore) FetchAllOwnersOfAProject(pid string) ([]modals.UserProjectMap, error) {
	var ownersmap []modals.UserProjectMap
	res := p.DB.Where("role = ?", constants.Role["owner"]).Where("pid = ?", pid).Find(&ownersmap)
	if res.RowsAffected == 0 {
		return ownersmap, errors.New("no owners found")
	}
	return ownersmap, res.Error
}
