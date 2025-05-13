package usuario

import (
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	dbtx          *dto.DbTransaction
	tokenBuilder  security.Builder
	tokenDuration time.Duration
}

func NewHandler(dbtx *dto.DbTransaction, tokenBuilder security.Builder, tokenDuration time.Duration) *Handler {
	return &Handler{dbtx: dbtx,
		tokenBuilder:  tokenBuilder,
		tokenDuration: tokenDuration}
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
		"result":  result, 
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

// Login de usuario
type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type loginResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}
type userResponse struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
}

func (h *Handler) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := h.dbtx.GetUsuarioByCorreo(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Correo o contraseña incorrectos"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.Password != req.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autorizado"})
		return
	}
	accessToken, err := h.tokenBuilder.CreateToken(user.Username, req.Email, user.Rol, h.tokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := loginResponse{
		AccessToken: accessToken,
		User: userResponse{
			UserName: user.Username,
			Role:     user.Rol,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
