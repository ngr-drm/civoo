package routes

import (
	"github.com/gin-gonic/gin"
	"github/Re44e/civoo/http/handlers/simulator"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		simulatorGroup := main.Group("simulator")
		{
			simulatorGroup.POST("/", simulator.Create)
			simulatorGroup.GET("/", simulator.List)
			simulatorGroup.PUT("/", simulator.Update)
			simulatorGroup.DELETE("/:id", simulator.Delete)
		}
	}
	return router
}