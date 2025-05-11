package usuario

import (
	"ProyectoProgramadoI/dto"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction) {
	h := NewHandler(dbtx)
	rg.POST("/", h.CreateUsuario)
	rg.GET("/:username", h.GetUsuarioByUsername)
	rg.GET("/", h.GetAllUsuarios)
	rg.PUT("/:username", h.UpdateUsuario)
	rg.DELETE("/:username", h.DeleteUsuario)
}
