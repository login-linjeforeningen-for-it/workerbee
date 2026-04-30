package handlers

import (
	"net/http"
	"strconv"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ListStorageBuckets(c *gin.Context) {
	buckets, err := h.Services.Storage.ListBuckets(c.Request.Context())
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"buckets": buckets})
}

func (h *Handler) CreateStorageBucket(c *gin.Context) {
	var body struct {
		Bucket string `json:"bucket" validate:"required"`
	}
	if err := c.ShouldBindBodyWithJSON(&body); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, body, *h.Services.Validate) {
		return
	}

	if err := h.Services.Storage.CreateBucket(c.Request.Context(), body.Bucket); internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) DeleteStorageBucket(c *gin.Context) {
	bucket := c.Param("bucket")
	if bucket == "" {
		internal.HandleError(c, internal.ErrInvalid)
		return
	}

	if err := h.Services.Storage.DeleteBucket(c.Request.Context(), bucket); internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) ListStorageObjects(c *gin.Context) {
	bucket := c.Query("bucket")
	prefix := c.Query("prefix")
	if bucket == "" {
		internal.HandleError(c, internal.ErrInvalid)
		return
	}

	objects, err := h.Services.Storage.ListObjects(c.Request.Context(), bucket, prefix)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"objects": objects})
}

func (h *Handler) UploadStorageObject(c *gin.Context) {
	bucket := c.PostForm("bucket")
	key := c.PostForm("key")
	file, err := c.FormFile("file")
	if internal.HandleError(c, err) {
		return
	}

	if bucket == "" || key == "" {
		internal.HandleError(c, internal.ErrInvalid)
		return
	}

	if err := h.Services.Storage.PutObject(c.Request.Context(), bucket, key, file); internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) MoveStorageObject(c *gin.Context) {
	var body models.S3ObjectMoveRequest
	if err := c.ShouldBindBodyWithJSON(&body); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, body, *h.Services.Validate) {
		return
	}

	if err := h.Services.Storage.MoveObject(c.Request.Context(), body); internal.HandleError(c, err) {
		return
	}

	mode := body.Mode
	if mode != "copy" {
		mode = "move"
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "mode": mode})
}

func (h *Handler) DeleteStorageObject(c *gin.Context) {
	bucket := c.Query("bucket")
	key := c.Query("key")
	if bucket == "" || key == "" {
		internal.HandleError(c, internal.ErrInvalid)
		return
	}

	if err := h.Services.Storage.DeleteObject(c.Request.Context(), bucket, key); internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) DownloadStorageObject(c *gin.Context) {
	bucket := c.Query("bucket")
	key := c.Query("key")
	if bucket == "" || key == "" {
		internal.HandleError(c, internal.ErrInvalid)
		return
	}

	body, contentType, contentLength, err := h.Services.Storage.GetObject(c.Request.Context(), bucket, key)
	if internal.HandleError(c, err) {
		return
	}
	defer body.Close()

	if contentType == "" {
		contentType = "application/octet-stream"
	}

	c.Header("Content-Disposition", `attachment; filename="`+key+`"`)
	if contentLength > 0 {
		c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, body, nil)
}
