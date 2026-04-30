package handlers

import (
	"net/http"
	"time"
	"workerbee/config"
	"workerbee/internal"
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

func (h *Handler) CreateStorageProof(c *gin.Context) {
	if config.StorageProofToken == "" || c.GetHeader("X-Workerbee-Storage-Proof") != config.StorageProofToken {
		internal.HandleError(c, internal.ErrUnauthorized)
		return
	}

	key, err := h.Services.ImageService.UploadStorageProof(c.Request.Context())
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bucket": internal.BUCKET_NAME,
		"key":    key,
		"url":    "https://s3.login.no/" + internal.BUCKET_NAME + "/" + key,
	})
}
