package simulator

import (
	"github.com/gin-gonic/gin"
	"github/Re44e/civoo/domain/services/simulator"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
	"github/Re44e/civoo/infrastructure/storage/pg/repository"
)

func List(data *gin.Context){
	var model []entities.Simulator
	var injection repository.PgRepository

	result, apiError := simulator.List(model, injection)

	if apiError != nil {
		data.JSON(apiError.Code, gin.H{
			"debug": apiError.Message,
			"error": apiError.Error.Error(),
		})
		return
	}

	data.JSON(200, result)
}