package transfer

import (
	"ProyectoProgramadoI/dto"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	dbtx *dto.DbTransaction
}

func NewHandler(dbtx *dto.DbTransaction) *Handler {
	return &Handler{dbtx: dbtx}
}

// Crear Transfer
type createTransferRequest struct {
	IdTransfer int32  `json:"idTransfer" binding:"required"`
	Tipo       string `json:"tipo" binding:"required"`
	Capacidad  int32  `json:"capacidad" binding:"required"`
}

func (h *Handler) CreateTransfer(ctx *gin.Context) {
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateTransferParams{
		Idtransfer: req.IdTransfer,
		Tipo:       req.Tipo,
		Capacidad:  req.Capacidad,
	}

	transfer, err := h.dbtx.CreateTransfer(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Transfer creado exitosamente",
		"transfer": transfer,
	})
}

// Obtener Transfer por ID
type getTransferByIdRequest struct {
	IdTransfer int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) GetTransferById(ctx *gin.Context) {
	var req getTransferByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transfer, err := h.dbtx.GetTransferById(ctx, req.IdTransfer)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Transfer no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfer)
}

// Obtener todos los Transfers
func (h *Handler) GetAllTransfers(ctx *gin.Context) {
	transfers, err := h.dbtx.GetAllTransfers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, transfers)
}

// Actualizar Transfer
type updateTransferUri struct {
	IdTransfer int32 `uri:"id" binding:"required,min=1"`
}

type updateTransferRequest struct {
	Tipo      string `json:"tipo" binding:"required"`
	Capacidad int32  `json:"capacidad" binding:"required"`
}

func (h *Handler) UpdateTransfer(ctx *gin.Context) {
	var uri updateTransferUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateTransferParams{
		Tipo:       req.Tipo,
		Capacidad:  req.Capacidad,
		Idtransfer: uri.IdTransfer,
	}

	err := h.dbtx.UpdateTransfer(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer actualizado exitosamente"})
}

// Eliminar Transfer
type deleteTransferRequest struct {
	IdTransfer int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) DeleteTransfer(ctx *gin.Context) {
	var req deleteTransferRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := h.dbtx.DeleteTransfer(ctx, req.IdTransfer)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Transfer no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer eliminado exitosamente"})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
