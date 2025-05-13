package middleware

import (
	"ProyectoProgramadoI/security"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Función pública para que sea accesible desde otros paquetes
func AuthMiddleware(tokenBilder security.Builder) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("authorization")
		if len(authHeader) == 0 {
			err := errors.New("falta token de autorización")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("formato de token inválido")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		if strings.ToLower(fields[0]) != "bearer" {
			err := errors.New("tipo de autorización no soportado: 'bearer' requerido")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenBilder.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set("authorized", payload)
		ctx.Next()
	}
}

// Verifica el rol del usuario para garantizar el acceso solo a usuarios con el rol correcto
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el payload del contexto
		authorized, exists := ctx.Get("authorized")
		if !exists {
			err := errors.New("usuario no autenticado")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// Convertir el payload al tipo adecuado (en este caso, Payload)
		payload, ok := authorized.(*security.Payload)
		if !ok {
			err := errors.New("información del usuario no válida")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		// Verificar el rol del usuario
		if payload.Rol != requiredRole {
			err := errors.New("acceso denegado: rol insuficiente")
			ctx.AbortWithStatusJSON(http.StatusForbidden, errorResponse(err))
			return
		}

		// Si el rol es el correcto, permitir la ejecución de la ruta
		ctx.Next()
	}
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
