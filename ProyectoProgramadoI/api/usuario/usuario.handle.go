package usuario

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

type createUsuarioRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Rol       string `json:"rol" binding:"required"`
	Idpersona int32  `json:"idpersona" binding:"required"`
}

func (h *Handler) CreateUsuario(ctx *gin.Context) {
	var req createUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.CreateUsuarioParams{
		Username:  req.Username,
		Password:  req.Password,
		Rol:       req.Rol,
		Idpersona: req.Idpersona,
	}

	result, err := h.dbtx.CreateUsuario(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuario creado exitosamente",
		"result":  result, // podría ser affectedRows o info del insert
	})
}

// Actualizar contraseña de usuario
type updateUsuarioRequest struct {
	Password string `json:"password" binding:"required"`
}

type updateUsuarioUri struct {
	Username string `uri:"username" binding:"required"`
}

func (h *Handler) UpdateUsuario(ctx *gin.Context) {
	var uri updateUsuarioUri
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := dto.UpdateUsuarioParams{
		Password: req.Password,
		Username: uri.Username,
	}

	err := h.dbtx.UpdateUsuario(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}

// Buscar por username
type getUsuarioByUsernameRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (h *Handler) GetUsuarioByUsername(ctx *gin.Context) {
	var req getUsuarioByUsernameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	usuario, err := h.dbtx.GetUsuarioByUserName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

// Eliminar usuarios...
type deleteUsuarioRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (h *Handler) DeleteUsuario(ctx *gin.Context) {
	var req deleteUsuarioRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := h.dbtx.DeleteUsuario(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}

// Obtener todos los usuarios para admin
func (h *Handler) GetAllUsuarios(ctx *gin.Context) {
	usuarios, err := h.dbtx.GetAllUsuarios(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, usuarios)
}

// Function reutiizable
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
