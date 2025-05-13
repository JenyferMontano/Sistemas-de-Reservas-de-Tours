package factura

import (
	"ProyectoProgramadoI/api/middleware"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction, tokenBuilder security.Builder) {
	h := NewHandler(dbtx)
	adminGroup := rg.Group("/")
	adminGroup.Use(
		middleware.AuthMiddleware(tokenBuilder),
		middleware.RequireRole("admin"),
	)
	adminGroup.POST("/", h.CreateFactura)
	adminGroup.GET("/", h.GetAllFacturas)
	rg.GET("/get/:id", h.GetFacturaById)
	adminGroup.DELETE("/:id", h.DeleteFactura)
}
