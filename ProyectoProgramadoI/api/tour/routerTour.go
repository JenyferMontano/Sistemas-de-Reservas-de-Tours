package tour

import (
	"ProyectoProgramadoI/dto"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction) {
	h := NewHandler(dbtx)
	rg.POST("/", h.CreateTour)
	rg.GET("/get/:id", h.GetTourById)
	rg.GET("/tipo/:tipo", h.GetToursByTipo)
	rg.GET("/", h.GetAllTours)
	rg.DELETE("/:id", h.DeleteTour)
	rg.PUT("/:id", h.UpdateTour)
}
