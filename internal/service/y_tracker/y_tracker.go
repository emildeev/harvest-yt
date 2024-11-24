package ytracker

import (
	"context"
	"time"

	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

type adapterI interface {
	GetMyself(ctx context.Context) (myself ytrackercore.User, err error)
	GetTicket(ctx context.Context, ticketKey string) (ticket ytrackercore.Ticket, err error)
	GetWorkLogs(ctx context.Context, createdBy string, from time.Time, to time.Time) (ytrackercore.WorkLogs, error)
	AddWorkLogs(
		ctx context.Context,
		taskKey string,
		start time.Time,
		duration time.Duration,
		comment string,
	) (ytrackercore.WorkLog, error)
}

type Service struct {
	adapter adapterI
}

func New(adapter adapterI) *Service {
	return &Service{
		adapter: adapter,
	}
}
