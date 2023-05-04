package service

import (
	"fmt"
	"strconv"

	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"golang.org/x/net/context"
)

type Racing interface {
	// ListRaces will return a collection of races.
	ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error)
	// GetRaceById returns race found by id
	GetRaceById(ctx context.Context, in *racing.GetRaceByIdRequest) (*racing.GetRaceByIdResponse, error)
}

// racingService implements the Racing interface.
type racingService struct {
	racesRepo db.RacesRepo
}

// NewRacingService instantiates and returns a new racingService.
func NewRacingService(racesRepo db.RacesRepo) Racing {
	return &racingService{racesRepo}
}

func (s *racingService) ListRaces(ctx context.Context, in *racing.ListRacesRequest) (*racing.ListRacesResponse, error) {
	races, err := s.racesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &racing.ListRacesResponse{Races: races}, nil
}

func (s *racingService) GetRaceById(ctx context.Context, in *racing.GetRaceByIdRequest) (*racing.GetRaceByIdResponse, error) {
	id, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil || id == 0 {
		return nil, fmt.Errorf("invalid Id given")
	}

	race, err := s.racesRepo.GetRaceById(id)
	if err != nil {
		return nil, err
	}

	return &racing.GetRaceByIdResponse{Race: race}, nil
}
