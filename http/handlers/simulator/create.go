package simulator

import (
	"github.com/gin-gonic/gin"
	"github/Re44e/civoo/domain/services/simulator"
	"github/Re44e/civoo/infrastructure/storage/pg/entities"
	"github/Re44e/civoo/infrastructure/storage/pg/repository"
)

func Create(data *gin.Context){
	var model entities.Simulator
	var injection repository.PgRepository
	err := data.ShouldBindJSON(&model)

	if err != nil {
		data.JSON(400, gin.H{
			"error": "Conversion to JSON failed: " + err.Error(),
		})
		return
	}

	result, apiError := simulator.Create(&model, injection)

	if apiError != nil {
		data.JSON(apiError.Code, gin.H{
			"debug": apiError.Message,
			"error": apiError.Error.Error(),
		})
		return
	}

	data.JSON(201, result)
}