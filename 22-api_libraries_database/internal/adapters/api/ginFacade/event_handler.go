package ginFacade

import (
	"net/http"
	"strconv"
	"time"

	"example.com/api/internal/core/domain"
	"example.com/api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	eventService ports.EventService
}

func NewEventHandler(service ports.EventService) *EventHandler {
	return &EventHandler{eventService: service}
}

func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.eventService.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	events, err := h.eventService.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var requestBody struct {
		Name        string    `binding:"required" json:"name"`
		Description string    `binding:"required" json:"description"`
		Location    string    `binding:"required" json:"location"`
		DateTime    time.Time `binding:"required" json:"date_time"`
		UserID      int       `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := &domain.Event{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		Location:    requestBody.Location,
		DateTime:    requestBody.DateTime,
		UserID:      requestBody.UserID,
	}

	createdEvent, err := h.eventService.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	c.JSON(http.StatusCreated, createdEvent)
}

func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	_, err = h.eventService.GetEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch event"})
		return
	}

	var updatedEvent domain.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse requested data"})
		return
	}
	updatedEvent.ID = id

	updatedEvent, err = h.eventService.UpdateEvent(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		return
	}
	c.JSON(http.StatusOK, updatedEvent)
}
