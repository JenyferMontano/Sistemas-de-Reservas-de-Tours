package transfer

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
	rg.POST("/", h.CreateTransfer)
	rg.GET("/:id", h.GetTransferById)
	rg.GET("/", h.GetAllTransfers)
	rg.PUT("/:id", h.UpdateTransfer)
	rg.DELETE("/:id", h.DeleteTransfer)
}
