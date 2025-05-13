package reserva

import (
	"ProyectoProgramadoI/dto"
	"database/sql"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	dbtx *dto.DbTransaction
}

// Constructor del handler
func NewHandler(dbtx *dto.DbTransaction) *Handler {
	return &Handler{dbtx: dbtx}
}

// Estructura para la creación de reserva con factura
type createReservaRequest struct {
	Numreserva       int32  `json:"numreserva"`
	Fechareserva     string `json:"fechareserva" binding:"required"`
	Horareserva      string `json:"horareserva" binding:"required"`
	Cantidadpersonas int32  `json:"cantidadpersonas" binding:"required"`
	Tour             int32  `json:"tour" binding:"required"`
	Usuario          string `json:"usuario" binding:"required"`
	Persona          int32  `json:"persona" binding:"required"`
	Transfer         int32  `json:"transfer" binding:"required"`
	Idfactura        int32  `json:"idfactura"`
	Metodopago       string `json:"metodopago" binding:"required"`
	Iva              string `json:"iva" binding:"required"`
	Descuento        string `json:"descuento" binding:"required"`
	Subtotal         string `json:"subtotal" binding:"required"`
	Total            string `json:"total" binding:"required"`
}

// Crear Reserva y Factura
func (h *Handler) CreateReserva(ctx *gin.Context) {
	var req createReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateReservaParams{
		Numreserva:       req.Numreserva,
		Fechareserva:     req.Fechareserva,
		Horareserva:      req.Horareserva,
		Cantidadpersonas: req.Cantidadpersonas,
		Tour:             req.Tour,
		Usuario:          req.Usuario,
		Persona:          req.Persona,
		Transfer:         req.Transfer,
	}

	// Ejecutar la creación de la reserva
	result, err := h.dbtx.CreateReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Obtener el ID de la reserva creada
	reservaID, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Crear la factura automáticamente después de la reserva
	facturaArgs := dto.CreateFacturaParams{
		Idfactura:  req.Idfactura,
		Fechafact:  req.Fechareserva,
		Metodopago: req.Metodopago,
		Iva:        req.Iva,
		Descuento:  req.Descuento,
		Subtotal:   req.Subtotal,
		Total:      req.Total,
		Reserva:    int32(reservaID),
	}

	_, err = h.dbtx.CreateFactura(ctx, facturaArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Respuesta de éxito
	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva y factura creadas correctamente"})
}

// Obtener reserva por ID
func (h *Handler) GetReservaById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Obtener la reserva por ID
	reserva, err := h.dbtx.GetReservaById(ctx, int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Reserva no encontrada"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Obtener la factura asociada a la reserva
	factura, err := h.dbtx.GetFacturaById(ctx, int32(id))
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Crear la respuesta incluyendo reserva y factura (si existe)
	response := gin.H{
		"reserva": reserva,
	}
	if err == nil {
		response["factura"] = factura
	}

	ctx.JSON(http.StatusOK, response)
}

// Eliminar reserva y su factura asociada
func (h *Handler) DeleteReserva(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Verificar si existe una factura asociada a la reserva
	facturaId, err := h.dbtx.GetFacturaByReservaId(ctx, int32(id))
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Si existe una factura, eliminarla primero
	if err == nil {
		err = h.dbtx.DeleteFactura(ctx, facturaId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	// Eliminar la reserva después de eliminar la factura
	err = h.dbtx.DeleteReserva(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva y factura eliminadas correctamente"})
}

// Estructura para la actualización de reserva con factura
type updateReservaRequest struct {
	Numreserva       int32  `json:"numreserva"`
	Fechareserva     string `json:"fechareserva" binding:"required"`
	Horareserva      string `json:"horareserva" binding:"required"`
	Cantidadpersonas int32  `json:"cantidadpersonas" binding:"required"`
	Tour             int32  `json:"tour" binding:"required"`
	Usuario          string `json:"usuario" binding:"required"`
	Persona          int32  `json:"persona" binding:"required"`
	Transfer         int32  `json:"transfer" binding:"required"`
	Metodopago       string `json:"metodopago" binding:"required"`
	Iva              string `json:"iva" binding:"required"`
	Descuento        string `json:"descuento" binding:"required"`
	Subtotal         string `json:"subtotal" binding:"required"`
	Total            string `json:"total" binding:"required"`
}

// Actualizar reserva y factura asociada
func (h *Handler) UpdateReserva(ctx *gin.Context) {
	var req updateReservaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Actualizar la reserva
	args := dto.UpdateReservaParams{
		Numreserva:       req.Numreserva,
		Fechareserva:     req.Fechareserva,
		Horareserva:      req.Horareserva,
		Cantidadpersonas: req.Cantidadpersonas,
		Tour:             req.Tour,
		Usuario:          req.Usuario,
		Persona:          req.Persona,
		Transfer:         req.Transfer,
	}

	err := h.dbtx.UpdateReserva(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Obtener el ID de la factura asociada a la reserva
	factura, err := h.dbtx.GetFacturaById(ctx, req.Numreserva)
	if err != nil && err != sql.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Solo actualiza la factura si ya existe
	if err == nil {
		facturaArgs := dto.UpdateFacturaParams{
			Metodopago: req.Metodopago,
			Iva:        req.Iva,
			Descuento:  req.Descuento,
			Subtotal:   req.Subtotal,
			Total:      req.Total,
			Idfactura:  factura.Idfactura,
		}

		err = h.dbtx.UpdateFactura(ctx, facturaArgs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva y factura actualizadas correctamente"})
}

// Obtener todas las reservas
func (h *Handler) GetAllReservas(ctx *gin.Context) {
	// Consultar todas las reservas desde la base de datos
	reservas, err := h.dbtx.GetAllReservas(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Verificar si no hay reservas
	if len(reservas) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No hay reservas registradas"})
		return
	}

	// Crear la estructura de respuesta
	var response []gin.H
	for _, reserva := range reservas {
		// Obtener la factura asociada a la reserva
		factura, err := h.dbtx.GetFacturaById(ctx, reserva.Numreserva)
		if err != nil && err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// Crear un objeto para la reserva con o sin factura
		item := gin.H{
			"numreserva":       reserva.Numreserva,
			"fechareserva":     reserva.Fechareserva,
			"horareserva":      reserva.Horareserva,
			"cantidadpersonas": reserva.Cantidadpersonas,
			"tour":             reserva.Nombretour,
			"usuario":          reserva.Nombreusuario,
			"persona":          reserva.Nombrepersona,
			"transfer":         reserva.Idtransfer,
		}

		// Agregar factura si existe
		if err == nil {
			item["factura"] = gin.H{
				"idfactura":  factura.Idfactura,
				"fechafact":  factura.Fechafact,
				"metodopago": factura.Metodopago,
				"iva":        factura.Iva,
				"descuento":  factura.Descuento,
				"subtotal":   factura.Subtotal,
				"total":      factura.Total,
			}
		}

		response = append(response, item)
	}

	// Respuesta de éxito
	ctx.JSON(http.StatusOK, response)
}

// Función reutilizable para manejar errores
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
