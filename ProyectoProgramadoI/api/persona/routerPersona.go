package persona

import (
	"ProyectoProgramadoI/dto"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction) {
	h := NewHandler(dbtx)
	rg.POST("/", h.CreatePersona)
	rg.GET("/get/:id", h.GetPersonaById)
	rg.GET("/", h.GetAllPersonas)
	rg.DELETE("/:id", h.DeletePersona)
	rg.PUT("/:id", h.UpdatePersona)
}
