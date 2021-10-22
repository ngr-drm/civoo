package simulator

import (
	"github.com/google/uuid"
	"github/Re44e/civoo/domain/interfaces/simulator"
	"github/Re44e/civoo/http/helper"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
)

func Create(data *entities.Simulator, storage simulator.Repository) (*entities.Simulator, *helper.ApiError)  {
	result, err := storage.Create(data)
	if err != nil {
		return nil, helper.ErrorHandler(helper.ApiError{ Error: err, Message: "creation service failure..." })
	}
	return result, nil
}

func List(reference []entities.Simulator, storage simulator.Repository) ([]entities.Simulator, *helper.ApiError)  {
	result, err := storage.List(reference)
	if err != nil {
		return nil, helper.ErrorHandler(helper.ApiError{ Error: err, Message: "query service failure..." })
	}
	return result, nil
}

func Update(data *entities.Simulator, storage simulator.Repository) (*entities.Simulator, *helper.ApiError)  {
	result, err := storage.Update(data)
	if err != nil {
		return nil, helper.ErrorHandler(helper.ApiError{ Error: err, Message: "update service failure..." })
	}
	return result, nil
}

func Delete(id uuid.UUID, storage simulator.Repository) *helper.ApiError  {
	err := storage.Delete(id)
	if err != nil {
		return helper.ErrorHandler(helper.ApiError{ Error: err, Message: "deletion service failure..." })
	}
	return nil
}