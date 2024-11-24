package validator

import (
	"context"
)

type yTrackerServiceI interface {
	GetTicketTitle(ctx context.Context, ticketKey string) (string, error)
}

type harvestServiceI interface {
	GetTasksListMap(ctx context.Context) (map[string]struct{}, error)
}

type UseCase struct {
	yTrackerService yTrackerServiceI
	harvestService  harvestServiceI
}

func New(
	yTrackerService yTrackerServiceI,
	harvestService harvestServiceI,
) *UseCase {
	return &UseCase{
		yTrackerService: yTrackerService,
		harvestService:  harvestService,
	}
}
