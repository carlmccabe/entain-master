package service

import (
	"fmt"
	"strconv"

	"sport/db"
	"sport/proto/sport"

	"golang.org/x/net/context"
)

type Sport interface {
	// ListEvents will return a collection of events.
	ListEvents(ctx context.Context, in *sport.ListEventsRequest) (*sport.ListEventsResponse, error)
	// GetEventById returns event found by id
	GetEventById(ctx context.Context, in *sport.GetEventByIdRequest) (*sport.GetEventByIdResponse, error)
}

// sportService implements the Sport interface.
type sportService struct {
	eventsRepo db.EventsRepo
}

// NewSportService instantiates and returns a new sportService.
func NewSportService(eventsRepo db.EventsRepo) Sport {
	return &sportService{eventsRepo}
}

func (s *sportService) ListEvents(ctx context.Context, in *sport.ListEventsRequest) (*sport.ListEventsResponse, error) {
	events, err := s.eventsRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &sport.ListEventsResponse{Events: events}, nil
}

func (s *sportService) GetEventById(ctx context.Context, in *sport.GetEventByIdRequest) (*sport.GetEventByIdResponse, error) {
	id, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil || id == 0 {
		return nil, fmt.Errorf("invalid Id given")
	}

	event, err := s.eventsRepo.GetEventById(id)
	if err != nil {
		return nil, err
	}

	return &sport.GetEventByIdResponse{Event: event}, nil
}
