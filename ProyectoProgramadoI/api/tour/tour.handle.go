package tour

import (
	"ProyectoProgramadoI/dto"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	dbtx *dto.DbTransaction
}

// Constructor del handler
func NewHandler(dbtx *dto.DbTransaction) *Handler {
	return &Handler{dbtx: dbtx}
}

type createTourRequest struct {
	Idtour         int32  `json:"idtour" binding:"required"`
	Nombre         string `json:"nombre" binding:"required"`
	Descripcion    string `json:"descripcion" binding:"required"`
	Tipo           string `json:"tipo" binding:"required"`
	Disponibilidad int8   `json:"disponibilidad" binding:"required"`
	Preciobase     string `json:"preciobase" binding:"required"`
	Ubicacion      string `json:"ubicacion" binding:"required"`
}

func (h *Handler) CreateTour(ctx *gin.Context) {
	var req createTourRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateTourParams{
		Idtour:         req.Idtour,
		Nombre:         req.Nombre,
		Descripcion:    req.Descripcion,
		Tipo:           req.Tipo,
		Disponibilidad: req.Disponibilidad,
		Preciobase:     req.Preciobase,
		Ubicacion:      req.Ubicacion,
	}

	tour, err := h.dbtx.CreateTour(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//ctx.JSON(http.StatusOK, tour) ESTO SI NO QUIERO EL MJS DE EXITO EVALUAR
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tour creado exitosamente",
		"tour":    tour,
	})
}

// Buscar por Id de tour
type getTourByIdRequest struct {
	Idtour int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) GetTourById(ptx *gin.Context) {
	var req getTourByIdRequest
	if err := ptx.ShouldBindUri(&req); err != nil {
		ptx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	persona, err := h.dbtx.GetTourById(ptx, req.Idtour)
	if err != nil {
		if err == sql.ErrNoRows {
			ptx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ptx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ptx.JSON(http.StatusOK, persona)
}

// Actualizar tours
type updateTourRequest struct {
	Nombre         string `json:"nombre" binding:"required"`
	Descripcion    string `json:"descripcion" binding:"required"`
	Tipo           string `json:"tipo" binding:"required"`
	Disponibilidad int8   `json:"disponibilidad" binding:"required"`
	Preciobase     string `json:"preciobase" binding:"required"`
	Ubicacion      string `json:"ubicacion" binding:"required"`
}

type updateTourUri struct {
	Idtour int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) UpdateTour(ctx *gin.Context) {
	var uri updateTourUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateTourRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateTourParams{
		Nombre:         req.Nombre,
		Descripcion:    req.Descripcion,
		Tipo:           req.Tipo,
		Disponibilidad: req.Disponibilidad,
		Preciobase:     req.Preciobase,
		Ubicacion:      req.Ubicacion,
		Idtour:         uri.Idtour,
	}

	err := h.dbtx.UpdateTour(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tour actualizado exitosamente",
	})
}

// Eliminar tours...
type deleteTourRequest struct {
	Idtour int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) DeleteTour(ptx *gin.Context) {
	var req deleteTourRequest
	if err := ptx.ShouldBindUri(&req); err != nil {
		ptx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := h.dbtx.DeleteTour(ptx, req.Idtour)
	if err != nil {
		if err == sql.ErrNoRows {
			ptx.JSON(http.StatusNotFound, gin.H{
				"message": "Tour no encontrado",
			})
			return
		}
		ptx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ptx.JSON(http.StatusOK, gin.H{
		"message": "Tour eliminado exitosamente",
	})
}

// Obtener todos los tours
func (h *Handler) GetAllTours(ctx *gin.Context) {
	tours, err := h.dbtx.GetAllTours(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tours)
}

//RECORDAR HACER LA FUNCION PARA BUSCAR POR TIPO DE TOUR

// Function reutiizable
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
