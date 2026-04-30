package services

import (
	"workerbee/internal"
	"workerbee/repositories"

	"github.com/go-playground/validator/v10"
)

type Services struct {
	Audiences     *AudienceService
	Categories    *CategoryService
	Events        *EventService
	Locations     *LocationService
	Organizations *OrganizationService
	Jobs          *JobsService
	Questions     *QuestionService
	Rules         *RuleService
	Stats         *StatsService
	Submissions   *SubmissionService
	ImageService  *ImageService
	Validate      *validator.Validate
	Honey         *HoneyService
	Alerts        *AlertService
	Albums        *AlbumService
	Storage       *StorageService
	Calendar      *CalendarService
	Compressor    *CompressorService
	Quotes        *QuoteService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Audiences:     NewAudienceService(repos.Audiences),
		Categories:    NewCategoryService(repos.Categories),
		Events:        NewEventService(repos.Events),
		Jobs:          NewJobsService(repos.Jobs),
		Questions:     NewQuestionService(repos.Questions),
		Rules:         NewRuleService(repos.Rules),
		Stats:         NewStatsService(repos.Stats),
		Submissions:   NewSubmissionService(repos.Submissions),
		Locations:     NewLocationService(repos.Locations),
		Organizations: NewOrganizationService(repos.Organizations),
		ImageService:  NewImageService(repos.Images),
		Honey:         NewHoneyService(repos.Honey),
		Alerts:        NewAlertService(repos.Alerts),
		Albums:        NewAlbumService(repos.Albums),
		Storage:       NewStorageService(repos.Storage),
		Calendar:      NewCalendarService(repos.Calendar),
		Compressor:    NewCompressorService(repos.Images, repos.Albums),
		Quotes:        NewQuoteService(repos.Quotes),
		Validate:      internal.SetUpValidator(),
	}
}
