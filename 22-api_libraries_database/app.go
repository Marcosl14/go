package main

import (
	apiGin "example.com/api/internal/adapters/api/ginProxy"
	"example.com/api/internal/adapters/api/ginProxy/router"
	"example.com/api/internal/adapters/db/sqlite"
	"example.com/api/internal/core/services"
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

	eventRepo := sqlite.NewSQLiteRepository(db)
	eventService := services.NewEventService(eventRepo)
	eventHandler := apiGin.NewEventHandler(eventService)

	// Configurar el servidor Gin
	server := gin.Default()

	// Usar el router para las rutas de eventos
	router.EventRoutes(server, eventHandler)

	// Iniciar el servidor
	server.Run(":8080")
}
