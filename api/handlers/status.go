package handlers

import (
	"net/http"
	"time"
	"workerbee/config"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// GetStatus godoc
// @Summary      Get API status
// @Description  Returns API version and uptime.
// @Tags         status
// @Produce      json
// @Success      200  {object}  models.Status
// @Router       /api/v2/status [get]
func GetStatus(c *gin.Context) {
	status := models.Status{
		Version: "v2",
		Uptime:  time.Duration(time.Since(config.StartTime).Seconds()),
	}

	c.JSON(http.StatusOK, status)
}
