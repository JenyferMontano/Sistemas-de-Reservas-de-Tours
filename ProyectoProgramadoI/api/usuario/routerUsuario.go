package usuario

import (
	"ProyectoProgramadoI/api/middleware"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction, builder security.Builder, tokenDuration time.Duration) {
	h := NewHandler(dbtx, builder, tokenDuration)
	adminGroup := rg.Group("/")
	adminGroup.Use(
		middleware.AuthMiddleware(builder),
		middleware.RequireRole("admin"),
	)
	adminGroup.POST("/", h.CreateUsuario)
	adminGroup.GET("/", h.GetAllUsuarios)
	adminGroup.GET("/:username", h.GetUsuarioByUsername)
	adminGroup.DELETE("/:username", h.DeleteUsuario)

	rg.PUT("/:username", middleware.AuthMiddleware(builder), h.UpdateUsuario)
}
