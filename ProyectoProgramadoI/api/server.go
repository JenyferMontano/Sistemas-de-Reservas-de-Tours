package api

import (
	"ProyectoProgramadoI/api/persona"
	"ProyectoProgramadoI/api/tour"
	"ProyectoProgramadoI/dto"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbtx   *dto.DbTransaction
	router *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction) *Server {
	server := &Server{dbtx: dbtx}
	router := gin.Default()
	//RUTAS {ENDPOINTS} DEL API
	api := router.Group("/api/v1")
	persona.RegisterRoutes(api.Group("/persona"), dbtx)
	tour.RegisterRoutes(api.Group("/tour"), dbtx)
	//router.POST("api/v1/category", server.createCategory)
	//router.GET("api/v1/category/:id", server.getCategory)
	///FIN RUTAS///
	server.router = router
	return server
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
