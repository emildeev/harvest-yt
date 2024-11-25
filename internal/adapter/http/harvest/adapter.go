package harvest

import (
	"context"
	"net/http"

	"github.com/emildeev/go-harvest/harvest"
)

const (
	serviceName = "harvest"
)

type timesheetI interface {
	List(
		ctx context.Context,
		opt *harvest.TimeEntryListOptions,
	) (*harvest.TimeEntryList, *http.Response, error)
	UpdateTimeEntry(
		ctx context.Context,
		timeEntryID int64,
		data *harvest.TimeEntryUpdate,
	) (*harvest.TimeEntry, *http.Response, error)
	Get(ctx context.Context, timeEntryID int64) (*harvest.TimeEntry, *http.Response, error)
}

type tasksI interface {
	List(ctx context.Context, opt *harvest.TaskListOptions) (*harvest.TaskList, *http.Response, error)
}

type Adapter struct {
	timesheet timesheetI
	tasks     tasksI
}

func New(timesheet timesheetI, tasks tasksI) *Adapter {
	return &Adapter{
		timesheet: timesheet,
		tasks:     tasks,
	}
}
