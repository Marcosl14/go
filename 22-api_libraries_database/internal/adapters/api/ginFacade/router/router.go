package router

import (
	"database/sql"

	"example.com/api/internal/adapters/api/ginFacade"
	"example.com/api/internal/adapters/db/sqlite"
	"example.com/api/internal/core/services"
	"github.com/gin-gonic/gin"
)

// SetupEventRoutes configura las rutas relacionadas con los eventos
func EventRoutes(server *gin.Engine, db *sql.DB) {
	eventRepo := sqlite.NewSQLiteRepository(db)
	eventService := services.NewEventService(eventRepo)
	eventHandler := ginFacade.NewEventHandler(eventService)

	mainPath := "/events"

	server.POST(mainPath, eventHandler.CreateEvent)
	server.GET(mainPath, eventHandler.GetEvents)
	server.GET(mainPath+"/:id", eventHandler.GetEvent)
}
