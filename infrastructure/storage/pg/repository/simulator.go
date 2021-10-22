package repository

import (
	"github.com/google/uuid"
	"github/Re44e/civoo/infrastructure/storage/pg"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
)

type PgRepository struct {}

func (s PgRepository) Create (payload *entities.Simulator) (*entities.Simulator, error) {
	instance := pg.GetDatabase()

	err := instance.Create(&payload).Error
	if err != nil {
			return nil, err
	}
	return payload, nil
}

func (s PgRepository) List (reference []entities.Simulator) ([]entities.Simulator, error) {
	instance := pg.GetDatabase()
	err := instance.Find(&reference).Error
	if err != nil {
		return nil, err
	}
	return reference, nil
}

func (s PgRepository) Update (payload *entities.Simulator) (*entities.Simulator, error) {
	instance := pg.GetDatabase()
	err := instance.Save(&payload).Error
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (s PgRepository) Delete (id uuid.UUID) error {
	instance := pg.GetDatabase()
	err := instance.Delete(&entities.Simulator{}, id).Error
	if err != nil {
		return err
	}
	return nil
}