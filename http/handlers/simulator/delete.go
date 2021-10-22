package simulator

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github/Re44e/civoo/domain/services/simulator"
	"github/Re44e/civoo/infrastructure/storage/pg/repository"
)

func Delete(data *gin.Context){
	id := data.Param("id")
	var injection repository.PgRepository
  validId, err := uuid.Parse(id)

	if err != nil {
		data.JSON(400, gin.H{
			"error": "id has to be uuid type",
		})
		return
	}

	apiError := simulator.Delete(validId, injection)

	if apiError != nil {
		data.JSON(apiError.Code, gin.H{
			"debug": apiError.Message,
			"error": apiError.Error.Error(),
		})

		return
	}

	data.Status(204)
}