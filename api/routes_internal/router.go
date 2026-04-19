package routes_internal

import (
	"workerbee/config"
	"workerbee/handlers"
	"workerbee/internal"
	"workerbee/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Route(c *gin.Engine, h *handlers.Handler) {
	c.GET("/", h.RootHandler)
	c.GET("/api", h.RootHandler)
	v2 := c.Group(internal.BASE_PATH)
	{
		v2.GET("", h.RootHandler)
		v2.GET("/ping", handlers.PingHandler)
		v2.GET("/docs", handlers.GetDocs)
		v2.GET("/status", handlers.GetStatus)
		events := v2.Group("/events")
		{
			events.GET("/protected/:id", middleware.AuthMiddleware(), h.GetProtectedEvent)
			events.GET("/:id", h.GetEvent)
			events.GET("/all", middleware.AuthMiddleware(), h.GetEventNames)
			events.GET("/protected", middleware.AuthMiddleware(), h.GetProtectedEvents)
			events.GET("", h.GetEvents)
			events.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateEvent,
			)
			events.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateEvent,
			)
			events.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteEvent,
			)
			events.GET("/categories", h.GetEventCategories)
			events.GET("/audiences", h.GetEventAudiences)
			events.GET("/time", h.GetAllTimeTypes)
		}
		rules := v2.Group("/rules")
		{
			rules.GET("/:id", h.GetRule)
			rules.GET("/all", h.GetRuleNames)
			rules.GET("", h.GetRules)
			rules.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateRule,
			)
			rules.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateRule,
			)
			rules.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteRule,
			)
		}
		categories := v2.Group("/categories")
		{
			categories.GET("/:id", h.GetCategory)
			categories.GET("", h.GetCategories)
			categories.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateCategory,
			)
			categories.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateCategory,
			)
			categories.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteCategory,
			)
		}
		locations := v2.Group("/locations")
		{
			locations.GET("/:id", h.GetLocation)
			locations.GET("/all", h.GetLocationNames)
			locations.GET("", h.GetLocations)
			locations.POST("", middleware.AuthMiddleware(), h.CreateLocation)
			locations.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateLocation,
			)
			locations.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteLocation,
			)
			locations.GET("/types", h.GetAllLocationTypes)
		}
		organizations := v2.Group("/organizations")
		{
			organizations.GET("/:id", h.GetOrganization)
			organizations.GET("/all", h.GetOrganizationNames)
			organizations.GET("", h.GetOrganizations)
			organizations.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateOrganization,
			)
			organizations.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateOrganization,
			)
			organizations.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteOrganization,
			)
		}
		jobs := v2.Group("/jobs")
		{
			jobs.GET("/:id", h.GetJob)
			jobs.GET("/protected/:id", middleware.AuthMiddleware(), h.GetProtectedJob)
			jobs.GET("", h.GetJobs)
			jobs.GET("/protected", middleware.AuthMiddleware(), h.GetProtectedJobs)
			jobs.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateJob,
			)
			jobs.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateJob,
			)
			jobs.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteJob,
			)
			jobs.GET("/cities", h.GetCities)
			jobs.GET("/skills", h.GetJobSkills)
			types := jobs.Group("/types")
			{
				types.GET("", h.GetActiveJobTypes)
				types.GET("/all", h.GetAllJobTypes)
				types.GET("/:id", h.GetJobType)
				types.POST("", middleware.AuthMiddleware(), h.CreateJobType)
				types.PUT("/:id", middleware.AuthMiddleware(), h.UpdateJobType)
				types.DELETE("/:id", middleware.AuthMiddleware(), h.DeleteJobType)
			}

		}
		audiences := v2.Group("/audiences")
		{
			audiences.GET("/:id", h.GetAudience)
			audiences.GET("", h.GetAudiences)
			audiences.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateAudience,
			)
			audiences.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateAudience,
			)
			audiences.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteAudience,
			)
		}
		stats := v2.Group("/stats")
		{
			stats.GET("/yearly", h.GetYearlyStats)
			stats.GET("/categories", h.GetMostActiveCategories)
			stats.GET("/new-additions", h.GetNewAdditionsStats)
		}
		images := v2.Group("/images/:path")
		{
			images.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UploadImage,
			)
			images.GET("", middleware.AuthMiddleware(), h.GetImageURLs)
			images.DELETE(
				"/:imageName",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteImage,
			)
		}
		honey := v2.Group("/honeys")
		{
			honey.GET("/:id", h.GetHoney)
			honey.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateHoney,
			)
			honey.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateHoney,
			)
			honey.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteHoney,
			)

		}
		text := v2.Group("/text")
		{
			text.GET("", h.GetTextServices)
			service := text.Group("/:service")
			{
				service.GET("", h.GetAllPathsInService)
				content := service.Group("/:path")
				{
					content.GET("", h.GetAllContentInPath)
					content.GET("/:language", h.GetOneLanguage)
				}
			}

		}
		alerts := v2.Group("/alerts")
		{
			alerts.GET("", h.GetAllAlerts)
			alerts.GET("/:service", h.GetAlertByServiceAndPage)
			alerts.GET("/id/:id", h.GetAlertByID)
			alerts.POST(
				"",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateAlert,
			)
			alerts.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateAlert,
			)
			alerts.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteAlert,
			)
		}
		albums := v2.Group("/albums")
		{
			albums.POST("", middleware.AuthMiddleware(), h.CreateAlbum)
			albums.POST("/:id", middleware.AuthMiddleware(), h.UploadImagesToAlbum)
			albums.GET("", h.GetAlbums)
			albums.GET("/:id", h.GetAlbum)
			albums.PUT(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateAlbum,
			)
			albums.DELETE(
				"/:id",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteAlbum,
			)
			albums.DELETE(
				"/:id/:imageName",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteAlbumImage,
			)
			albums.PUT(
				"/:id/:imageName",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.SetAlbumCover,
			)
			albums.PUT(
				"/compress",
				middleware.AuthMiddleware(),
				middleware.RateLimitMiddleware(1),
				h.CompressAlbumImages,
			)
		}
		calendar := v2.Group("/calendar")
		{
			calendar.GET("", h.GetCalendar)
		}
		quotes := v2.Group("/quotes")
		{
			quotes.POST(
				"",
				middleware.QuoteMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.CreateQuote,
			)
			quotes.GET("", h.GetQuotes)
			quotes.DELETE(
				"/:id",
				middleware.QuoteMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.DeleteQuote,
			)
			quotes.PUT(
				"/:id",
				middleware.QuoteMiddleware(),
				middleware.RateLimitMiddleware(config.AllowedRequestsPerMinute),
				h.UpdateQuote,
			)
		}
	}
}
