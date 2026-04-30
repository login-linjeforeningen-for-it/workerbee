package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrConflict               = errors.New("resource already exists")
	ErrNoImagesProvided       = errors.New("no images provided")
	ErrNotFound               = errors.New("could not find id")
	ErrNoRow                  = errors.New("no row found")
	ErrUnableToDelete         = errors.New("unable to delete resource")
	ErrInvalid                = errors.New("invalid user data")
	ErrInvalidImagePath       = errors.New("invalid image path")
	ErrImageTooLarge          = errors.New("image size exceeds maximum limit")
	ErrInvalidImageRatio      = errors.New("invalid image aspect ratio")
	ErrUnauthorized           = errors.New("unauthorized opperation")
	ErrInvalidForeignKey      = errors.New("error foreign key does not exist")
	ErrInvalidAudience        = errors.New("invalid audience does not exist in enum")
	ErrInvalidTimeType        = errors.New("invalid time type does not exist in enum")
	ErrInvalidCategory        = errors.New("invalid category does not exist in enum")
	ErrInvalidLocationType    = errors.New("invalid location type does not exist in enum")
	ErrInvalidJobType         = errors.New("invalid job type does not exist in enum")
	ErrTooManyRequests        = errors.New("too many requests")
	ErrCacheUnavailable       = errors.New("cache unavailable")
	ErrUnknownImageFormat     = errors.New("unknown image format")
	ErrS3ClientNotInitialized = errors.New("s3 client not initialized")
	ErrorMap                  = map[error]struct {
		Status  int
		Message string
	}{
		ErrTooManyRequests:   {Status: http.StatusTooManyRequests, Message: "too many requests"},
		ErrNoImagesProvided:  {Status: http.StatusBadRequest, Message: "no images provided"},
		ErrNotFound:          {Status: http.StatusBadRequest, Message: "did not find document"},
		ErrNoRow:             {Status: http.StatusBadRequest, Message: "no row found"},
		ErrInvalid:           {Status: http.StatusBadRequest, Message: "invalid user data"},
		ErrInvalidImagePath:  {Status: http.StatusBadRequest, Message: "invalid image path"},
		ErrUnauthorized:      {Status: http.StatusUnauthorized, Message: "unauthorized operation"},
		ErrInvalidForeignKey: {Status: http.StatusBadRequest, Message: "error foreign key does not exist"},
		ErrInvalidAudience:   {Status: http.StatusBadRequest, Message: "invalid audience does not exist in enum"},
		ErrInvalidTimeType:   {Status: http.StatusBadRequest, Message: "invalid time type does not exist in enum"},
		ErrInvalidCategory:   {Status: http.StatusBadRequest, Message: "invalid category does not exist in enum"},
		ErrInvalidLocationType: {
			Status:  http.StatusBadRequest,
			Message: "invalid location type does not exist in enum",
		},
		ErrInvalidJobType:         {Status: http.StatusBadRequest, Message: "invalid job type does not exist in enum"},
		ErrImageTooLarge:          {Status: http.StatusBadRequest, Message: "image size exceeds maximum limit, max 1MB"},
		ErrInvalidImageRatio:      {Status: http.StatusBadRequest, Message: "invalid image aspect ratio, max 2.5"},
		ErrS3ClientNotInitialized: {Status: http.StatusInternalServerError, Message: "s3 client not initialized"},
		ErrConflict:               {Status: http.StatusConflict, Message: "resource already exists"},
		ErrUnknownImageFormat:     {Status: http.StatusBadRequest, Message: "unknown image format please use jpg or png"},
		ErrUnableToDelete:         {Status: http.StatusBadRequest, Message: "unable to delete resource, either too old or invalid ID"},
	}
)

func logRequestError(c *gin.Context, status int, err error) {
	method := ""
	path := ""
	route := ""
	query := ""
	remoteAddr := ""
	userAgent := ""
	params := gin.Params(nil)

	if c != nil && c.Request != nil {
		method = c.Request.Method
		path = c.Request.URL.Path
		query = c.Request.URL.RawQuery
		remoteAddr = c.ClientIP()
		userAgent = c.Request.UserAgent()
		route = c.FullPath()
		params = c.Params
	}

	log.Printf(
		"Got error: err=%v status=%d method=%s path=%s route=%s query=%q remote=%s user_agent=%q params=%v",
		err,
		status,
		method,
		path,
		route,
		query,
		remoteAddr,
		userAgent,
		params,
	)
}

func HandleError(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	for k, v := range ErrorMap {
		if errors.Is(err, k) {
			c.JSON(v.Status, gin.H{"error": v.Message})
			logRequestError(c, v.Status, err)
			return true
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	logRequestError(c, http.StatusInternalServerError, err)
	return true
}

func HandleValidationError[T any](c *gin.Context, body T, validate validator.Validate) bool {
	if err := validate.Struct(body); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationMessages := make([]string, 0, len(errs))
			eventType := reflect.TypeOf(body)
			if eventType.Kind() == reflect.Ptr {
				eventType = eventType.Elem()
			}

			for _, e := range errs {
				fieldName := e.Field()
				if f, found := eventType.FieldByName(e.StructField()); found {
					if jsonTag := f.Tag.Get("json"); jsonTag != "" {
						fieldName = strings.Split(jsonTag, ",")[0]
					}
				}
				validationMessages = append(validationMessages, fmt.Sprintf("%s failed on %s validation", fieldName, e.Tag()))
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationMessages,
			})
			return true
		}
		HandleError(c, err)
		return true
	}
	return false
}
