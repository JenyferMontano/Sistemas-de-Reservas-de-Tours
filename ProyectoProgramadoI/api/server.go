package api

import (
	"ProyectoProgramadoI/api/persona"
	"ProyectoProgramadoI/api/tour"
	"ProyectoProgramadoI/api/transfer"
	"ProyectoProgramadoI/api/usuario"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbtx          *dto.DbTransaction
	tokenBuilder  security.Builder
	tokenDuration time.Duration
	router        *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction, tokenDuration time.Duration) (*Server, error) {
	//server := &Server{dbtx: dbtx}
	tokenBuilder, err := security.NewPasetoBuilder("12345678123456781234567812345678")
	if err != nil {
		return nil, err
	}
	server := &Server{
		dbtx:          dbtx,
		tokenBuilder:  tokenBuilder,
		tokenDuration: tokenDuration,
	}
	router := gin.Default()
	usuarioHandler := usuario.NewHandler(dbtx, tokenBuilder, tokenDuration)

	//RUTAS {ENDPOINTS} DEL API
	api := router.Group("/api/v1")
	api.POST("/login", usuarioHandler.Login)
	persona.RegisterRoutes(api.Group("/persona"), dbtx, tokenBuilder)
	tour.RegisterRoutes(api.Group("/tour"), dbtx, tokenBuilder)
	usuario.RegisterRoutes(api.Group("/usuario"), dbtx, tokenBuilder, tokenDuration)
	transfer.RegisterRoutes(api.Group("/transfer"), dbtx, tokenBuilder)

	///FIN RUTAS///
	server.router = router
	return server, nil
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
