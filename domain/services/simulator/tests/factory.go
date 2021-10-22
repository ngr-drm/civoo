package tests

import (
	"github.com/google/uuid"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
)

func MockCreate () entities.Simulator {
	id, _ := uuid.Parse("3bdbc864-51bd-4a6f-a7e9-667970793947")
	data := entities.Simulator{
	 	ID: id,
		Date: "27/07/2021",
		Landing: 2,
		Qualification: "IFRH",
		Role: "Piloto",
	}
	return data
}