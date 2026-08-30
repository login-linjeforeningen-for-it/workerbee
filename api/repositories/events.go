package repositories

import (
	"errors"
	"os"
	"time"
	"workerbee/db"
	"workerbee/internal"
	"workerbee/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Eventrepositories interface {
	GetProtectedEvents(limit, offset int, search, orderBy, sort string, historical bool, categories, audiences []int) ([]models.EventWithTotalCount, error)
	GetEvents(limit, offset int, search, orderBy, sort string, categories, audiences []int) ([]models.EventWithTotalCount, error)
	GetEvent(id string) (models.Event, error)
	GetProtectedEvent(id string) (models.Event, error)
	GetEventCategories() ([]models.EventCategory, error)
	DeleteEvent(id string) (int, error)
	UpdateOneEvent(id int, event models.NewEvent) (models.NewEvent, error)
	CreateEvent(event models.NewEvent) (models.NewEvent, error)
	CreateMultipleEvents(event models.NewEvent, repeatUntil internal.Date, repeatType string) (models.NewEvent, error)
	GetEventAudiences() ([]models.Audience, error)
	GetAllTimeTypes() (string, error)
	GetEventNames() ([]models.EventName, error)
	GetNextPublishTime() (*time.Time, error)
}

type eventRepositories struct {
	db *sqlx.DB
}

func NewEventrepositories(db *sqlx.DB) Eventrepositories {
	return &eventRepositories{db: db}
}

func (r *eventRepositories) CreateEvent(event models.NewEvent) (models.NewEvent, error) {
	return db.AddOneRow(
		r.db,
		"./db/events/post_event.sql",
		event,
	)
}

func (r *eventRepositories) CreateMultipleEvents(event models.NewEvent, repeatUntil internal.Date, repeatType string) (models.NewEvent, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return models.NewEvent{}, err
	}
	defer tx.Rollback()

	sqlBytes, err := os.ReadFile("./db/events/create_event_return_id.sql")
	if err != nil {
		return models.NewEvent{}, err
	}
	stmt, err := tx.PrepareNamed(string(sqlBytes))
	if err != nil {
		return models.NewEvent{}, err
	}
	defer stmt.Close()

	start := event.TimeStart.Time
	end := event.TimeEnd.Time

	var signupRelease *time.Time
	var signupDeadline *time.Time

	if event.TimeSignupRelease != nil {
		signupRelease = &event.TimeSignupRelease.Time
	}

	if event.TimeSignupDeadline != nil {
		signupDeadline = &event.TimeSignupDeadline.Time
	}

	// Ensure repeatUntil.Time is set to the end of the day (23:59:59)
	repeatUntilEndOfDay := time.Date(
		repeatUntil.Time.Year(),
		repeatUntil.Time.Month(),
		repeatUntil.Time.Day(),
		23, 59, 59, 0,
		repeatUntil.Time.Location(),
	)

	for {
		if start.After(repeatUntilEndOfDay) {
			break
		}

		newEvent := event
		newEvent.TimeStart = internal.LocalTime{Time: start}
		newEvent.TimeEnd = internal.LocalTime{Time: end}

		if signupRelease != nil {
			newEvent.TimeSignupRelease = &internal.LocalTime{Time: *signupRelease}
		} else {
			newEvent.TimeSignupRelease = nil
		}

		if signupDeadline != nil {
			newEvent.TimeSignupDeadline = &internal.LocalTime{Time: *signupDeadline}
		} else {
			newEvent.TimeSignupDeadline = nil
		}

		var id int
		rows, err := stmt.Queryx(newEvent)
		if err != nil {
			return models.NewEvent{}, err
		}
		if rows.Next() {
			if err := rows.Scan(&id); err != nil {
				return models.NewEvent{}, err
			}
		}
		rows.Close()

		var deltaDays, offset int
		switch repeatType {
		case "weekly":
			deltaDays = 7
			offset = 0
		case "biweekly":
			deltaDays = 14
			offset = 7
		default:
			return models.NewEvent{}, internal.ErrInvalid
		}

		// The next signup and deadlines are shifted to be right after the previous event and before the next event
		if signupRelease != nil {
			*signupRelease = end.AddDate(0, 0, 1+offset) // ✅ correct (def not ai generated)
		}

		start = start.AddDate(0, 0, deltaDays)
		end = end.AddDate(0, 0, deltaDays)

		if signupDeadline != nil {
			*signupDeadline = (*signupDeadline).AddDate(0, 0, deltaDays) // ✅ correct (def not ai generated)
		}
	}

	if err := tx.Commit(); err != nil {
		return models.NewEvent{}, err
	}
	return event, nil
}

func (r *eventRepositories) GetProtectedEvents(limit, offset int, search, orderBy, sort string, historical bool, categories, audiences []int) ([]models.EventWithTotalCount, error) {
	events, err := db.FetchAllElements[models.EventWithTotalCount](
		r.db,
		"./db/events/get_protected_events.sql",
		orderBy, sort,
		limit,
		offset,
		search,
		historical,
		pq.Array(categories),
		pq.Array(audiences),
	)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepositories) GetEvents(limit, offset int, search, orderBy, sort string, categories, audiences []int) ([]models.EventWithTotalCount, error) {
	events, err := db.FetchAllElements[models.EventWithTotalCount](
		r.db,
		"./db/events/get_events.sql",
		orderBy, sort,
		limit,
		offset,
		search,
		pq.Array(categories),
		pq.Array(audiences),
	)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepositories) GetEventCategories() ([]models.EventCategory, error) {

	categories, err := db.FetchAllForeignAttributes[models.EventCategory](
		r.db,
		"./db/events/get_categories.sql",
	)

	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *eventRepositories) GetEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/get_event.sql", id)
	if err != nil {
		if errors.Is(err, internal.ErrNoRow) {
			return models.Event{}, internal.ErrNotFound
		}
		return models.Event{}, internal.ErrInvalid
	}
	return event, nil
}

func (r *eventRepositories) GetProtectedEvent(id string) (models.Event, error) {
	event, err := db.ExecuteOneRow[models.Event](r.db, "./db/events/get_protected_event.sql", id)
	if err != nil {
		if errors.Is(err, internal.ErrNoRow) {
			return models.Event{}, internal.ErrNotFound
		}
		return models.Event{}, internal.ErrInvalid
	}
	return event, nil
}

func (r *eventRepositories) UpdateOneEvent(id int, event models.NewEvent) (models.NewEvent, error) {
	event.ID = id

	return db.AddOneRow(r.db, "./db/events/put_event.sql", event)
}

func (r *eventRepositories) DeleteEvent(id string) (int, error) {
	eventId, err := db.DeleteOneRow[int](r.db, "./db/events/delete_event.sql", id)
	if err != nil {
		return 0, err
	}
	return eventId, nil
}

func (r *eventRepositories) GetEventAudiences() ([]models.Audience, error) {
	audiences, err := db.FetchAllForeignAttributes[models.Audience](
		r.db,
		"./db/events/get_event_audiences.sql",
	)
	if err != nil {
		return nil, err
	}
	return audiences, nil
}

func (r *eventRepositories) GetAllTimeTypes() (string, error) {
	timeTypes, err := db.FetchAllEnumTypes(r.db, "./db/events/get_all_time_types.sql")
	if err != nil {
		return "", err
	}
	return timeTypes, nil
}

func (r *eventRepositories) GetEventNames() ([]models.EventName, error) {
	eventNames, err := db.FetchAllForeignAttributes[models.EventName](
		r.db,
		"./db/events/get_event_names.sql",
	)
	if err != nil {
		return nil, err
	}
	return eventNames, nil
}

func (r *eventRepositories) GetNextPublishTime() (*time.Time, error) {
	nextPublishTime, err := db.ExecuteOneRow[*time.Time](r.db, "./db/events/get_next_publish_time.sql")
	if err != nil {
		return nil, err
	}
	return nextPublishTime, nil
}
