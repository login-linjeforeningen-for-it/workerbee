package repositories

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Audiences     Audiencerepository
	Categories    Categoryrepository
	Locations     LocationRepository
	Organizations OrganizationRepository
	Events        Eventrepositories
	Jobs          Jobsrepositories
	Questions     Questionrepositories
	Rules         Rulerepositories
	Stats         Statsrepositories
	Submissions   Submissionrepositories
	Honey         HoneyRepository
	Alerts        AlertRepository
	Albums        AlbumsRepository
	Images        ImageRepository
	Calendar      CalendarRepository
	Quotes        QuoteRepository
}

func NewRepositories(db *sqlx.DB, objectStorage *s3.Client) *Repositories {
	return &Repositories{
		Audiences:     NewAudiencerepository(db),
		Categories:    NewCategoryRepository(db),
		Events:        NewEventrepositories(db),
		Jobs:          NewJobrepositories(db),
		Questions:     NewQuestionrepositories(db),
		Rules:         NewRulerepositories(db),
		Stats:         NewStatsrepositories(db),
		Submissions:   NewSubmissionrepositories(db),
		Locations:     NewLocationRepository(db),
		Organizations: NewOrganizationRepository(db),
		Honey:         NewHoneyRepository(db),
		Alerts:        NewAlertRepository(db),
		Albums:        NewAlbumsRepository(db, objectStorage),
		Images:        NewImageRepository(db, objectStorage),
		Calendar:      NewCalendarRepository(db),
		Quotes:        NewQuoteRepository(db),
	}
}
