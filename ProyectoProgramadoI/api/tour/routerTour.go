package tour

import (
	"ProyectoProgramadoI/api/middleware"
	"ProyectoProgramadoI/dto"

	"ProyectoProgramadoI/security"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction, tokenBuilder security.Builder) {
	h := NewHandler(dbtx)
	rg.Use(
		middleware.AuthMiddleware(tokenBuilder),
		middleware.RequireRole("admin"),
	)
	rg.POST("/", h.CreateTour)
	rg.GET("/get/:id", h.GetTourById)
	rg.GET("/tipo/:tipo", h.GetToursByTipo)
	rg.GET("/", h.GetAllTours)
	rg.DELETE("/:id", h.DeleteTour)
	rg.PUT("/:id", h.UpdateTour)
}
