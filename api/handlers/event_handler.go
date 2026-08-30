package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateEvent godoc
// @Summary      Create a new event
// @Description  Creates a new event with the provided details. Requires authentication.
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event  body      models.NewEvent  true  "Event to create"
// @Success      201    {object}  models.Event
// @Failure      400    {object}  error
// @Failure      500    {object}  error
// @Router       /api/v2/events [post]
func (h *Handler) CreateEvent(c *gin.Context) {
	var event models.NewEvent

	repeatUntil := c.DefaultQuery("repeat_until", "")
	repeatType := c.DefaultQuery("repeat_type", "")

	if err := c.ShouldBindBodyWithJSON(&event); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, event, *h.Services.Validate) {
		return
	}

	event, err := h.Services.Events.CreateEvent(event, repeatUntil, repeatType)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, event)
}

// UpdateEvent godoc
// @Summary      Update an existing event
// @Description  Updates an existing event with the provided details. Requires authentication.
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        id     path      string          true  "Event ID"
// @Param        event  body      models.NewEvent  true  "Updated event details"
// @Success      200    {object}  models.Event
// @Failure      400    {object}  error
// @Failure      500    {object}  error
// @Router       /api/v2/events/{id} [put]
func (h *Handler) UpdateEvent(c *gin.Context) {
	var event models.NewEvent
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&event); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, event, *h.Services.Validate) {
		return
	}

	event, err := h.Services.Events.UpdateEvent(event, id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, event)
}

// GetEvents godoc
// @Summary      Get events
// @Description  Returns a list of events with details, including category, location, audience, and organizer info. Supports historical filtering, limit, and offset.
// @Tags         events
// @Produce      json
// @Param        limit    query  int   false  "Maximum number of results"
// @Param        offset   query  int   false  "Offset for pagination"
// @Param        historical query bool false  "Include historical events"
// @Success      200  {array}  models.Event
// @Failure      500  {object}  error
// @Router       /api/v2/events [get]
func (h *Handler) GetEvents(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	categories := c.DefaultQuery("categories", "")
	audiences := c.DefaultQuery("audiences", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "time_start")
	sort := c.DefaultQuery("sort", "asc")

	events, cacheTTL, err := h.Services.Events.GetEvents(search, limit, offset, orderBy, sort, categories, audiences)
	if internal.HandleError(c, err) {
		return
	}

	c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", cacheTTL))

	if len(events) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": events[0].TotalCount,
		})
	}
}

// GetEvents godoc
// @Summary      Get events
// @Description  Returns a list of events with details, including category, location, audience, and organizer info. Supports historical filtering, limit, and offset.
// @Tags         events
// @Produce      json
// @Param        limit    query  int   false  "Maximum number of results"
// @Param        offset   query  int   false  "Offset for pagination"
// @Param        historical query bool false  "Include historical events"
// @Success      200  {array}  models.Event
// @Failure      500  {object}  error
// @Router       /api/v2/events [get]
func (h *Handler) GetProtectedEvents(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	categories := c.DefaultQuery("categories", "")
	audiences := c.DefaultQuery("audiences", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")
	historical := c.DefaultQuery("historical", "false")

	events, err := h.Services.Events.GetProtectedEvents(search, limit, offset, orderBy, sort, historical, categories, audiences)
	if internal.HandleError(c, err) {
		return
	}
	if len(events) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"events":      events,
			"total_count": events[0].TotalCount,
		})
	}
}

// GetEvent godoc
// @Summary      Get event by ID
// @Description  Retrieves a specific event by its ID, including detailed information such as category, location, audience, and organizer info.
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event ID"
// @Success      200  {object}  models.Event
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/events/{id} [get]
func (h *Handler) GetEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.Services.Events.GetEvent(id)
	if errors.Is(err, internal.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, event)
}

// GetProtectedEvent godoc
// @Summary      Get protected event by ID
// @Description  Retrieves a specific protected event by its ID, including detailed information such as category, location, audience, and organizer info.
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event ID"
// @Success      200  {object}  models.Event
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/events/{id} [get]
func (h *Handler) GetProtectedEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := h.Services.Events.GetProtectedEvent(id)
	if errors.Is(err, internal.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, event)
}

// GetEventAudiences godoc
// @Summary      Get event audiences
// @Description  Retrieves a list of all event audiences.
// @Tags         events
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/events/audiences [get]
func (h *Handler) GetEventAudiences(c *gin.Context) {
	audiences, err := h.Services.Events.GetEventAudiences()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, audiences)
}

// GetAllTimeTypes godoc
// @Summary      Get all time types
// @Description  Retrieves a list of all time types for events.
// @Tags         events
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/events/time [get]
func (h *Handler) GetAllTimeTypes(c *gin.Context) {
	timeTypes, err := h.Services.Events.GetAllTimeTypes()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, timeTypes)
}

// GetEventCategories godoc
// @Summary      Get event categories
// @Description  Retrieves a list of all event categories active.
// @Tags         events
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/events/categories [get]
func (h *Handler) GetEventCategories(c *gin.Context) {
	categories, err := h.Services.Events.GetEventCategories()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, categories)
}

// DeleteEvent godoc
// @Summary      Delete an event
// @Description  Deletes an event by its ID. Requires authentication.
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/events/{id} [delete]
func (h *Handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	eventId, err := h.Services.Events.DeleteEvent(id)
	if internal.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": eventId})
}

func (h *Handler) GetEventNames(c *gin.Context) {
	eventNames, err := h.Services.Events.GetEventNames()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, eventNames)
}
