package db

import (
	"gihtub.com/heyyakash/switchr/internal/modals"
	"gorm.io/gorm"
)

func (p *PostgresStore) CreateProject(project *modals.Projects) error {
	result := p.DB.Create(&project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *PostgresStore) GetAllProjectsByUid(uid string) ([]modals.Projects, error) {
	var projects []modals.Projects
	res := p.DB.Where("uid = ?", uid).Find(&projects)
	if res.Error != nil {
		return nil, res.Error
	}
	return projects, nil
}
func (p *PostgresStore) GetProjectByPid(pid string) (modals.Projects, error) {
	var projects modals.Projects
	res := p.DB.Where("pid = ?", pid).First(&projects)
	return projects, res.Error
}

func (p *PostgresStore) DeleteProject(project *modals.Projects) error {
	res := p.DB.Delete(&project).Error
	return res
}
func (p *PostgresStore) DeleteProjectById(id uint) error {
	res := p.DB.Where("id = ?", id).Delete(&modals.Projects{}).Error
	return res
}
func (p *PostgresStore) DeleteProjectByPid(pid string) error {
	return p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("pid = ?", pid).Delete(&modals.UserProjectMap{}).Error; err != nil {
			return err
		}
		if err := tx.Where("pid = ?", pid).Delete(&modals.Projects{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (p *PostgresStore) GetProjectById(id uint) (modals.Projects, error) {
	var project modals.Projects
	res := p.DB.Where("id = ?", id).First(&project)
	return project, res.Error
}

func (p *PostgresStore) GetFlagByPid(pid string) ([]modals.FeatureflagWithUserName, error) {
	var flags []modals.FeatureflagWithUserName
	// res := p.DB.Where("pid = ?", pid).Find(&flags)
	res := p.DB.Table("featureflags").
		Select("featureflags.*, users.full_name").
		Joins("JOIN users ON users.uid = featureflags.created_by").
		Where("featureflags.pid = ?::uuid", pid).
		Find(&flags)

	return flags, res.Error
}

func (p *PostgresStore) UpdateProjectWithPid(project *modals.Projects, pid string) error {
	return p.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(modals.Projects{}).Where("pid = ?", pid).Updates(project).Error; err != nil {
			return err
		}
		return nil
	})
}
