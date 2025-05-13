package factura

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

//Crear factura

type createFacturaRequest struct {
	Idfactura  int32  `json:"idfactura"`
	Fechafact  string `json:"fechafact"`
	Metodopago string `json:"metodopago"`
	Iva        string `json:"iva"`
	Descuento  string `json:"descuento"`
	Subtotal   string `json:"subtotal"`
	Total      string `json:"total"`
	Reserva    int32  `json:"reserva"`
}

func (h *Handler) CreateFactura(ctx *gin.Context) {
	var req createFacturaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateFacturaParams{
		Idfactura:  req.Idfactura,
		Fechafact:  req.Fechafact,
		Metodopago: req.Metodopago,
		Iva:        req.Iva,
		Descuento:  req.Descuento,
		Subtotal:   req.Subtotal,
		Total:      req.Total,
		Reserva:    req.Reserva,
	}

	factura, err := h.dbtx.CreateFactura(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, factura)
}

// Buscar facura por id
type getFacturaByIdRequest struct {
	Idfactura int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) GetFacturaById(ctx *gin.Context) {
	var req getFacturaByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	factura, err := h.dbtx.GetFacturaById(ctx, req.Idfactura)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, factura)
}

// Obtener todas las facturas
func (h *Handler) GetAllFacturas(ctx *gin.Context) {
	facturas, err := h.dbtx.GetAllFacturas(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, facturas)
}

// Eliminar Factura
func (h *Handler) DeleteFactura(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.dbtx.DeleteFactura(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Factura eliminada correctamente"})
}

// Function reutiizable
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
