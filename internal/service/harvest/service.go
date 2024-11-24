package harvest

import (
	"context"
	"regexp"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
)

type adapterI interface {
	GetTimesheetListToday(ctx context.Context) (harvestcore.TimeEntries, error)
	GetTimesheet(ctx context.Context, id int64) (harvestcore.TimeEntry, error)
	UpdateTimesheetComment(ctx context.Context, id int64, comment string) (harvestcore.TimeEntry, error)
	GetAllTasks(ctx context.Context) ([]string, error)
}

type Service struct {
	adapter adapterI

	pushedRegexp *regexp.Regexp
	tasksCache   []string
}

const (
	pushedEmoji = "\u2705"
	pushRegexp  = pushedEmoji + "([0-9]+)"
	pushFormat  = pushedEmoji + "(%d)"
)

func New(adapter adapterI) *Service {
	return &Service{
		adapter: adapter,

		pushedRegexp: regexp.MustCompile(pushRegexp),
	}
}