package persona

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
	rg.POST("/", h.CreatePersona)
	rg.GET("/get/:id", h.GetPersonaById)
	rg.GET("/", h.GetAllPersonas)
	rg.DELETE("/:id", h.DeletePersona)
	rg.PUT("/:id", h.UpdatePersona)
}
