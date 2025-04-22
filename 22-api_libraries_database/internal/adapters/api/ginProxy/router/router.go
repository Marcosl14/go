package router

import (
	"github.com/gin-gonic/gin"

	"example.com/api/internal/adapters/api/ginProxy"
)

// SetupEventRoutes configura las rutas relacionadas con los eventos
func EventRoutes(router *gin.Engine, eventHandler *ginProxy.EventHandler) {
	router.GET("/events", eventHandler.GetEvents)
	router.POST("/events", eventHandler.CreateEvent)
	// Aquí puedes agregar más rutas relacionadas con eventos
	// como GET /events/:id, PUT /events/:id, DELETE /events/:id, etc.
}
