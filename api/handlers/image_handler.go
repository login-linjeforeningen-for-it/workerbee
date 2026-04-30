package handlers

import (
	"net/http"
	"path/filepath"
	"strings"
	"workerbee/internal"

	"github.com/gin-gonic/gin"

	_ "image/jpeg"
)

// UploadImage godoc
// @Summary      Upload an image to a specified path in object storage
// @Description  Uploads an image file to the specified path in the image service.
// @Tags         images
// @Accept       multipart/form-data
// @Produce      json
// @Param        path   path      string  true  "Path to upload the image"
// @Param        image  formData  file    true  "Image file to upload"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  error
// @Failure      500    {object}  error
// @Router       /api/v2/images/{path} [post]
func (h *Handler) UploadImage(c *gin.Context) {
	path := c.Param("path")
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, internal.MaxImageUploadSize)

	if c.Request.ContentLength > internal.MaxImageUploadSize {
		internal.HandleError(c, internal.ErrImageTooLarge)
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		if strings.Contains(err.Error(), "request body too large") {
			internal.HandleError(c, internal.ErrImageTooLarge)
			return
		}
		internal.HandleError(c, err)
		return
	}

	imageURL, err := h.Services.ImageService.UploadImage(file, c.Request.Context(), path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"image": imageURL, "name": filepath.Base(imageURL)})
}

// GetImageURLs godoc
// @Summary      Get image URLs in a specified path in object storage
// @Description  Retrieves a list of image URLs available in the specified path.
// @Tags         images
// @Produce      json
// @Param        path   path      string  true  "Path to retrieve images from"
// @Success      200    {array}   string
// @Failure      500    {object}  error
// @Router       /api/v2/images/{path} [get]
func (h *Handler) GetImageURLs(c *gin.Context) {
	path := c.Param("path")

	imageURLs, err := h.Services.ImageService.GetImagesInPath(c.Request.Context(), path)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, imageURLs)
}

// DeleteImage godoc
// @Summary      Delete an image from a specified path in object storage
// @Description  Deletes an image file from the specified path in the image service.
// @Tags         images
// @Produce      json
// @Param        path       path      string  true  "Path where the image is located"
// @Param        imageName  path      string  true  "Name of the image to delete"
// @Success      200        {object}  map[string]string
// @Failure      400        {object}  error
// @Failure      500        {object}  error
// @Router       /api/v2/images/{path}/{imageName} [delete]
func (h *Handler) DeleteImage(c *gin.Context) {
	path := c.Param("path")
	imageName := c.Param("imageName")

	path, err := h.Services.ImageService.DeleteImage(c.Request.Context(), path, imageName)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, path)
}
