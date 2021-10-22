package tests

import (
	"github.com/google/uuid"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
	"gorm.io/gorm"
	"os/exec"
	"time"
)

type MockRepository struct {}

func (s MockRepository) Create (payload *entities.Simulator) (*entities.Simulator, error) {
	var simulator []*entities.Simulator
	if payload != nil {
		created := append(simulator, payload)
		return created[0], nil
	}
	return nil, new(exec.Error)
}

func (s MockRepository) List (reference []entities.Simulator) ([]entities.Simulator, error) {
	if reference != nil {
		mockQuery := append(reference, MockCreate())
		return mockQuery, nil
	}
	return nil, new(exec.Error)
}

func (s MockRepository) Update (payload *entities.Simulator) (*entities.Simulator, error) {
	if payload != nil {
		mockData := MockCreate()
		mockCurrentState := make([]entities.Simulator, 1)
		mockUpdate := make([]entities.Simulator, 1)

		mockCurrentState[0] = mockData

		payload.UpdatedAt = time.Now()
		mockUpdate[0] = *payload
		copy(mockCurrentState, mockUpdate)

		return &mockCurrentState[0], nil
	}
	return nil, new(exec.Error)
}

func (s MockRepository) Delete(id uuid.UUID) error {
	mockData := MockCreate()
	mockCurrentState := make([]entities.Simulator, 1)
	mockCurrentState[0] = mockData

	if mockCurrentState[0].ID == id {
		mockCurrentState[0].DeletedAt = gorm.DeletedAt{Time: time.Now()}
		return nil
	}

	return new(exec.Error)

}

