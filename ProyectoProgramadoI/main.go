package main

import (
	"ProyectoProgramadoI/api"
	"ProyectoProgramadoI/dto"
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	dbDriver  = "mysql"
	dbSource  = "root:@tcp(127.0.0.1:3306)/reservas?charset=utf8mb4&parseTime=True&loc=Local"
	serverURL = "127.0.0.1:8080"
)

func main() {

	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("No se pudo cargar el archivo .env")
	}

	tokenDurationStr := os.Getenv("TOKEN_DURATION")
	tokenDuration, err := time.ParseDuration(tokenDurationStr)
	if err != nil {
		log.Fatal("Duración del token inválida:", err)
	}

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("No se puede establecer la conexión", err)
	}
	dbtx := dto.NewDbTransaction(conn)
	server, err := api.NewServer(dbtx, tokenDuration)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
	err = server.Start(serverURL)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
}
