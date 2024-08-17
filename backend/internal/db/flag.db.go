package db

import (
	"gihtub.com/heyyakash/switchr/internal/modals"
)

func (p *PostgresStore) CreateFlag(flag *modals.Featureflag) error {
	err := p.DB.Create(&flag).Error
	return err
}

func (p *PostgresStore) GetFlagsByProjectId(pid string) ([]modals.Featureflag, error) {
	var flags []modals.Featureflag
	res := p.DB.Where("project_id = ?", pid).Find(&flags)
	return flags, res.Error
}

func (p *PostgresStore) UpdateFlag(id uint, value string) error {
	err := p.DB.Where("id = ?").Update("value", value).Error
	return err
}

func (p *PostgresStore) DeleteFlag(id uint) error {
	err := p.DB.Where("id = ?", id).Delete(&modals.Featureflag{}).Error
	return err
}

func (p *PostgresStore) GetFlagById(id uint) (modals.Featureflag, error) {
	var flag modals.Featureflag
	res := p.DB.Where("id = ?", id).First(&flag)
	return flag, res.Error
}

func (p *PostgresStore) GetFlagByFid(fid string) (modals.Featureflag, error) {
	var flag modals.Featureflag
	res := p.DB.Where("fid = ?", fid).First(&flag)
	return flag, res.Error
}
