package harvest

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/emildeev/harvest-yt/internal/adapter/http/harvest/port"
	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
	httpcore "github.com/emildeev/harvest-yt/internal/core/http"
)

func (adapter *Adapter) GetTimesheetListToday(ctx context.Context) (harvestcore.TimeEntries, error) {
	handleError := httpcore.GetHandleErrorFunc[harvestcore.TimeEntries](
		serviceName, "UpdateTimeEntry", nil,
	)

	list, resp, err := adapter.timesheet.List(ctx, port.GetTodayOpt())
	response := &resty.Response{
		RawResponse: resp,
	}
	if err != nil {
		return handleError(err, response)
	}

	if err := httpcore.HandleHTTPError(err, response); err != nil {
		return handleError(err, response)
	}

	return port.GetTimeEntriesFormTimeEntryList(list), nil
}

func (adapter *Adapter) GetTimesheet(ctx context.Context, id int64) (harvestcore.TimeEntry, error) {
	handleError := httpcore.GetHandleErrorFunc[harvestcore.TimeEntry](
		serviceName, "UpdateTimeEntry", harvestcore.TimeEntry{},
	)

	timer, resp, err := adapter.timesheet.Get(ctx, id)
	response := &resty.Response{
		RawResponse: resp,
	}
	if err != nil {
		return handleError(err, response)
	}

	if err := httpcore.HandleHTTPError(err, response); err != nil {
		return handleError(err, response)
	}

	return port.GetTimeEntryFromTimeEntry(timer), nil
}

func (adapter *Adapter) UpdateTimesheetComment(
	ctx context.Context,
	id int64,
	comment string,
) (harvestcore.TimeEntry, error) {
	handleError := httpcore.GetHandleErrorFunc[harvestcore.TimeEntry](
		serviceName, "UpdateTimeEntry", harvestcore.TimeEntry{},
	)

	timeEntry, resp, err := adapter.timesheet.UpdateTimeEntry(ctx, id, port.GetTimeEntryUpdateCommentRequest(comment))
	response := &resty.Response{
		RawResponse: resp,
	}
	if err != nil {
		return handleError(err, response)
	}

	if err := httpcore.HandleHTTPError(err, response); err != nil {
		return handleError(err, response)
	}

	return port.GetTimeEntryFromTimeEntry(timeEntry), nil
}
