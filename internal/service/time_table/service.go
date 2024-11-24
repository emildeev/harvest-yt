package timetable

import (
	"context"
	"regexp"

	"github.com/emildeev/harvest-yt/internal/config"
	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

const (
	taskKeyRegexp = ytrackercore.TicketKeyRegexp + ":"
)

type (
	yTrackerServiceI interface {
		ValidateTicketForSpend(ctx context.Context, taskKey string) (ytrackercore.Ticket, error)
	}
)

type Service struct {
	yTrackerService yTrackerServiceI

	cfg           config.Tasks
	taskKeyRegexp *regexp.Regexp
}

func New(
	yTrackerService yTrackerServiceI,

	cfg config.Tasks,
) *Service {
	return &Service{
		yTrackerService: yTrackerService,

		cfg:           cfg,
		taskKeyRegexp: regexp.MustCompile(taskKeyRegexp),
	}
}
