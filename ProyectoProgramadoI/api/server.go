package api

import (
	"ProyectoProgramadoI/api/factura"
	"ProyectoProgramadoI/api/persona"
	"ProyectoProgramadoI/api/reserva"
	"ProyectoProgramadoI/api/tour"
	"ProyectoProgramadoI/api/transfer"
	"ProyectoProgramadoI/api/usuario"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/security"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type Server struct {
	dbtx          *dto.DbTransaction
	tokenBuilder  security.Builder
	tokenDuration time.Duration
	router        *gin.Engine
}

func NewServer(dbtx *dto.DbTransaction, tokenDuration time.Duration) (*Server, error) {
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
	// Middleware CORS
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*", 
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	usuarioHandler := usuario.NewHandler(dbtx, tokenBuilder, tokenDuration)

	//RUTAS {ENDPOINTS} DEL API
	api := router.Group("/api/v1")
	api.POST("/login", usuarioHandler.Login)
	persona.RegisterRoutes(api.Group("/persona"), dbtx, tokenBuilder)
	tour.RegisterRoutes(api.Group("/tour"), dbtx, tokenBuilder)
	usuario.RegisterRoutes(api.Group("/usuario"), dbtx, tokenBuilder, tokenDuration)
	transfer.RegisterRoutes(api.Group("/transfer"), dbtx, tokenBuilder)
	reserva.RegisterRoutes(api.Group("/reserva"), dbtx, tokenBuilder)
	factura.RegisterRoutes(api.Group("/factura"), dbtx, tokenBuilder)

	///FIN RUTAS///
	server.router = router
	return server, nil
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}
