package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"sport/proto/sport"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)

// EventsRepo provides repository access to events.
type EventsRepo interface {
	// Init will initialise our events repository.
	Init() error

	// List will return a list of events.
	List(filter *sport.ListEventsRequestFilter) ([]*sport.Event, error)

	// GetEventById
	GetEventById(id int64) (*sport.Event, error)
}

type eventsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewEventsRepo creates a new events repository.
func NewEventsRepo(db *sql.DB) EventsRepo {
	return &eventsRepo{db: db}
}

// Init prepares the events repository dummy data.
func (r *eventsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy events.
		err = r.seed()
	})

	return err
}

func (r *eventsRepo) List(filter *sport.ListEventsRequestFilter) ([]*sport.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventQueries()[eventsList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanEvents(rows)
}

func (r *eventsRepo) GetEventById(id int64) (*sport.Event, error) {
	query := getEventQueries()[event]
	row := r.db.QueryRow(query, strconv.FormatInt(id, 10))

	var event sport.Event
	var advertisedStart time.Time

	err := row.Scan(&event.Id, &event.Home, &event.Away, &event.Number, &event.Visible, &advertisedStart)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	ts, err := ptypes.TimestampProto(advertisedStart)
	if err != nil {
		return nil, err
	}

	event.Status = getStatus(advertisedStart)
	event.AdvertisedStartTime = ts

	return &event, nil
}

func (r *eventsRepo) applyFilter(query string, filter *sport.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	if filter.Visible {
		clauses = append(clauses, "visible = ?")
		args = append(args, true)
	}

	if filter.Hidden {
		clauses = append(clauses, "visible = ?")
		args = append(args, false)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	if filter.OrderBy != nil {
		var sort string = fmt.Sprintf("ORDER BY %s %s", filter.OrderBy.Field, getSortDirection(filter.OrderBy.Ascending))
		query += " " + sort
	}

	return query, args
}

func (m *eventsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sport.Event, error) {
	var events []*sport.Event

	for rows.Next() {
		var event sport.Event
		var advertisedStart time.Time

		if err := rows.Scan(&event.Id, &event.Home, &event.Away, &event.Number, &event.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.Status = getStatus(advertisedStart)
		event.AdvertisedStartTime = ts

		events = append(events, &event)
	}

	return events, nil
}

// Helper function to determine the sort direction
func getSortDirection(ascending bool) string {
	if ascending {
		return "ASC"
	}
	return "DESC"
}

// Helper function to determine status
func getStatus(startTime time.Time) string {
	currentTime := time.Now().Local()
	if currentTime.After(startTime) {
		return "closed"
	}

	return "open"
}
