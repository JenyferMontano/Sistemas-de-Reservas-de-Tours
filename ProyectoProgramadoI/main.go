package main

import (
	"ProyectoProgramadoI/api"
	"ProyectoProgramadoI/dto"
	"ProyectoProgramadoI/utils"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("No se pudo cargar el archivo de configuración:", err)
	}

	tokenDuration, err := time.ParseDuration(config.TokenDuration)
	if err != nil {
		log.Fatal("Duración del token inválida:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("No se puede establecer la conexión", err)
	}
	dbtx := dto.NewDbTransaction(conn)
	server, err := api.NewServer(dbtx, tokenDuration)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
	err = server.Start(config.ServerURL)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
}
