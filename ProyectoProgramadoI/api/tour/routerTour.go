package tour

import (
	"ProyectoProgramadoI/api/middleware"
	"ProyectoProgramadoI/dto"

	"ProyectoProgramadoI/security"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction, tokenBuilder security.Builder) {
	h := NewHandler(dbtx)
	// Aplicar middlewares directamente al grupo principal `tour`
	rg.Use(
		middleware.AuthMiddleware(tokenBuilder), // Verifica el token
		middleware.RequireRole("admin"),         // Verifica que sea admin
	)
	rg.POST("/", h.CreateTour)
	rg.GET("/get/:id", h.GetTourById)
	rg.GET("/tipo/:tipo", h.GetToursByTipo)
	rg.GET("/", h.GetAllTours)
	rg.DELETE("/:id", h.DeleteTour)
	rg.PUT("/:id", h.UpdateTour)
}
