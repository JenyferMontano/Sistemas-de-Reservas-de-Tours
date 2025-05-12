package api

import (
	"ProyectoProgramadoI/api/persona"
	"ProyectoProgramadoI/api/tour"
	"ProyectoProgramadoI/api/transfer"
	"ProyectoProgramadoI/api/usuario"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbtx         *dto.DbTransaction
	tokenBuilder security.Builder
	router       *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction) (*Server, error) {
	//server := &Server{dbtx: dbtx}
	tokenBuilder, err := security.NewPasetoBuilder("12345678123456781234567812345678")
	if err != nil {
		return nil, err
	}
	server := &Server{
		dbtx:         dbtx,
		tokenBuilder: tokenBuilder,
	}
	router := gin.Default()
	//RUTAS {ENDPOINTS} DEL API
	api := router.Group("/api/v1")
	persona.RegisterRoutes(api.Group("/persona"), dbtx)
	tour.RegisterRoutes(api.Group("/tour"), dbtx)
	usuario.RegisterRoutes(api.Group("/usuario"), dbtx)
	transfer.RegisterRoutes(api.Group("/transfer"), dbtx)
	//router.POST("api/v1/category", server.createCategory)
	//router.GET("api/v1/category/:id", server.getCategory)

	//RUTAS CON MIDDLEWARE

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
