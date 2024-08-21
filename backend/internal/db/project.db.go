package db

import "gihtub.com/heyyakash/switchr/internal/modals"

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
	res := p.DB.Where("pid = ?", pid).Delete(&modals.Projects{}).Error
	return res
}

func (p *PostgresStore) GetProjectById(id uint) (modals.Projects, error) {
	var project modals.Projects
	res := p.DB.Where("id = ?", id).First(&project)
	return project, res.Error
}

func (p *PostgresStore) GetFlagByPid(pid string) ([]modals.Featureflag, error) {
	var flags []modals.Featureflag
	res := p.DB.Where("pid = ?", pid).Find(&flags)
	return flags, res.Error
}
