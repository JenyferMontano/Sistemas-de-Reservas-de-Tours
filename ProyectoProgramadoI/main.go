package main

import (
	"ProyectoProgramadoI/api"
	"ProyectoProgramadoI/dto"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver  = "mysql"
	dbSource  = "root:@tcp(127.0.0.1:3306)/reservas?charset=utf8mb4&parseTime=True&loc=Local"
	serverURL = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("No se puede establecer la conexión", err)
	}
	dbtx := dto.NewDbTransaction(conn)
	server := api.NewServer(dbtx)
	err = server.Start(serverURL)
	if err != nil {
		log.Fatal("No se puede iniciar el servidor", err)
	}
}
