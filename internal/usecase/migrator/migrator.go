package migrator

import (
	"context"
	"time"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
	timetablecore "github.com/emildeev/harvest-yt/internal/core/time_table"
	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

type (
	yTrackerServiceI interface {
		GetTodayTimeTable(ctx context.Context) (ytrackercore.WorkLogs, error)
		SpendTime(
			ctx context.Context,
			taskKey string,
			startTime time.Time,
			duration time.Duration,
			comment string,
		) (int, error)
	}
	harvestServiceI interface {
		GetTimersList(
			ctx context.Context,
		) (pushed harvestcore.TimeEntries, notPushed harvestcore.TimeEntries, err error)
		MarkTimerPushed(ctx context.Context, id int64, workloadId int) error
	}
	timeTableI interface {
		Generate(ctx context.Context, entries harvestcore.TimeEntries, offset time.Duration) timetablecore.Table
	}
)

type UseCase struct {
	yTrackerService yTrackerServiceI
	harvestService  harvestServiceI
	timeTable       timeTableI
}

func New(
	yTrackerService yTrackerServiceI,
	harvestService harvestServiceI,
	timeTable timeTableI,
) *UseCase {
	return &UseCase{
		yTrackerService: yTrackerService,
		harvestService:  harvestService,
		timeTable:       timeTable,
	}
}
