package simulator

import (
	"github.com/google/uuid"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
)

type Repository interface {
	Create(payload *entities.Simulator) (*entities.Simulator, error)
	List(reference []entities.Simulator) ([]entities.Simulator, error)
	Update(payload *entities.Simulator) (*entities.Simulator, error)
	Delete(id uuid.UUID)  error
}
