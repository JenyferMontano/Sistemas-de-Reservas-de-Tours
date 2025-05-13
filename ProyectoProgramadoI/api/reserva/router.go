package reserva

import (
	"ProyectoProgramadoI/api/middleware"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, dbtx *dto.DbTransaction, tokenBuilder security.Builder) {
	h := NewHandler(dbtx)

	// Rutas para gestionar reservas
	rg.POST("/", h.CreateReserva)                                                                           // Crear una nueva reserva (incluye factura)
	rg.GET("/", middleware.AuthMiddleware(tokenBuilder), middleware.RequireRole("admin"), h.GetAllReservas) // Obtener todas las reservas
	rg.GET("/:id", h.GetReservaById)                                                                        // Obtener reserva por ID
	rg.PUT("/:id", h.UpdateReserva)                                                                         // Actualizar reserva (y factura si aplica)
	rg.DELETE("/:id", h.DeleteReserva)                                                                      // Eliminar reserva (y factura asociada)
}
