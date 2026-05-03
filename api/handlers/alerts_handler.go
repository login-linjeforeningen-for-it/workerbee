package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateAlert godoc
// @Summary      Create a new alert
// @Description  Creates a new alert with the provided details
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        alert body      models.Alert true "Alert details"
// @Success      201  {object}  models.Alert
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/alerts [post]
func (h *Handler) CreateAlert(c *gin.Context) {
	var alert models.Alert

	if err := c.ShouldBindBodyWithJSON(&alert); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, alert, *h.Services.Validate) {
		return
	}

	alertResponse, err := h.Services.Alerts.CreateAlert(alert)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, alertResponse)
}

// GetAllAlerts godoc
// @Summary      Get all alerts
// @Description  Retrieves a list of all alerts with optional search, pagination, and sorting
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Limit number of results"
// @Param        offset    query     string  false  "Offset for results"
// @Param        order_by  query     string  false  "Field to order by"
// @Param        sort      query     string  false  "Sort direction (asc or desc)"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  error
// @Router       /api/v2/alerts [get]
func (h *Handler) GetAllAlerts(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	alerts, err := h.Services.Alerts.GetAllAlerts(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(alerts) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"alerts":      alerts,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"alerts":      alerts,
			"total_count": alerts[0].TotalCount,
		})
	}
}

// GetAlertByServiceAndPage godoc
// @Summary      Get alert by service and page
// @Description  Retrieves an alert based on the specified service and page
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        service   path      string  true  "Service name"
// @Param        page      path      string  true  "Page name"
// @Success      200  {object}  models.Alert
// @Failure      400  {object}  error
// @Router       /api/v2/alerts/service/{service}/page/{page} [get]
func (h *Handler) GetAlertByServiceAndPage(c *gin.Context) {
	service := c.Param("service")

	page := c.DefaultQuery("page", "/")

	alert, err := h.Services.Alerts.GetAlertByServiceAndPage(service, page)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alert)
}

// GetAlertByID godoc
// @Summary      Get alert by ID
// @Description  Retrieves an alert based on the specified ID
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Alert ID"
// @Success      200  {object}  models.Alert
// @Failure      400  {object}  error
// @Router       /api/v2/alerts/{id} [get]
func (h *Handler) GetAlertByID(c *gin.Context) {
	id := c.Param("id")

	alert, err := h.Services.Alerts.GetAlertByID(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alert)
}

// UpdateAlert godoc
// @Summary      Update an alert
// @Description  Updates an existing alert with the provided details
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Alert ID"
// @Param        alert body      models.Alert true "Updated alert details"
// @Success      200  {object}  models.Alert
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/alerts/{id} [put]
func (h *Handler) UpdateAlert(c *gin.Context) {
	var alert models.Alert
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&alert); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, alert, *h.Services.Validate) {
		return
	}

	alertResponse, err := h.Services.Alerts.UpdateAlert(id, alert)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alertResponse)
}

// DeleteAlert godoc
// @Summary      Delete an alert
// @Description  Deletes an existing alert by ID
// @Tags         alerts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Alert ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Security     Bearer
// @Router       /api/v2/alerts/{id} [delete]
func (h *Handler) DeleteAlert(c *gin.Context) {
	id := c.Param("id")

	deletedID, err := h.Services.Alerts.DeleteAlert(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedID})
}
