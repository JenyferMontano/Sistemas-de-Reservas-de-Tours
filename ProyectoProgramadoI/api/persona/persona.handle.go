package persona

import (
	"ProyectoProgramadoI/dto"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	dbtx *dto.DbTransaction
}

func NewHandler(dbtx *dto.DbTransaction) *Handler {
	return &Handler{dbtx: dbtx}
}

type createPersonaRequest struct {
	Idpersona int32     `json:"id_persona" binding:"required"`
	Nombre    string    `json:"nombre" binding:"required"`
	Apellido1 string    `json:"apellido_1" binding:"required"`
	Apellido2 string    `json:"apellido_2" binding:"required"`
	FechaNac  time.Time `json:"fecha_nac" binding:"required"`
	Telefono  string    `json:"telefono" binding:"required"`
	Correo    string    `json:"correo" binding:"required,email"`
}

func (h *Handler) CreatePersona(ctx *gin.Context) {
	var req createPersonaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreatePersonaParams{
		Idpersona: req.Idpersona,
		Nombre:    req.Nombre,
		Apellido1: req.Apellido1,
		Apellido2: req.Apellido2,
		Fechanac:  req.FechaNac,
		Telefono:  req.Telefono,
		Correo:    req.Correo,
	}

	persona, err := h.dbtx.CreatePersona(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, persona)
}

// Obtener persona por id
type getPersonaByIdRequest struct {
	Idpersona int32 `uri:"id" binding:"required,min=1"`
}

func (h *Handler) GetPersonaById(ptx *gin.Context) {
	var req getPersonaByIdRequest
	if err := ptx.ShouldBindUri(&req); err != nil {
		ptx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	persona, err := h.dbtx.GetPersonaById(ptx, req.Idpersona)
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

// Listar Personas
func (h *Handler) GetAllPersonas(ctx *gin.Context) {
	personas, err := h.dbtx.GetAllPersonas(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, personas)
}

// Eliminar Persona
func (h *Handler) DeletePersona(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = h.dbtx.DeletePersona(ctx, int32(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Persona eliminada correctamente"})
}

// Actualizar Persona
func (h *Handler) UpdatePersona(ctx *gin.Context) {
	var uri struct {
		Idpersona int32 `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req createPersonaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdatePersonaParams{
		Idpersona: uri.Idpersona,
		Nombre:    req.Nombre,
		Apellido1: req.Apellido1,
		Apellido2: req.Apellido2,
		Fechanac:  req.FechaNac,
		Telefono:  req.Telefono,
		Correo:    req.Correo,
	}

	err := h.dbtx.UpdatePersona(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
