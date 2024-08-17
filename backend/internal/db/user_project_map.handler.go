package db

import "gihtub.com/heyyakash/switchr/internal/modals"

func (p *PostgresStore) CreateUserProjectMap(userprojectmap *modals.UserProjectMap) error {
	err := p.DB.Create(&userprojectmap).Error
	return err
}
