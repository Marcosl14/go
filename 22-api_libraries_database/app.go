package main

import (
	"example.com/api/internal/adapters/api/ginFacade/router"
	"example.com/api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	db, err := database.InitDB("api.db")

	if err != nil {
		panic("Error initializing database: " + err.Error())
	}

	defer db.Close()

	// Configurar el servidor Gin
	server := gin.Default()

	// Iniciar el router de eventos
	router.EventRoutes(server, db)

	// Iniciar el servidor
	server.Run(":8080")
}
