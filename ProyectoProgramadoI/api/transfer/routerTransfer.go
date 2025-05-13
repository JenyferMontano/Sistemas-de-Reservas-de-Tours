package transfer

import (
	"ProyectoProgramadoI/dto"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction) {
	h := NewHandler(dbtx)
	rg.POST("/", h.CreateTransfer)
	rg.GET("/:id", h.GetTransferById)
	rg.GET("/", h.GetAllTransfers)
	rg.PUT("/:id", h.UpdateTransfer)
	rg.DELETE("/:id", h.DeleteTransfer)
}
