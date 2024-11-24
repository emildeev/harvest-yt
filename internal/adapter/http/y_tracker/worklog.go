package ytracker

import (
	"context"
	"time"

	"github.com/emildeev/harvest-yt/internal/adapter/http/y_tracker/port"
	httpcore "github.com/emildeev/harvest-yt/internal/core/http"
	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func (adapter *Adapter) GetWorkLogs(
	_ context.Context,
	createdBy string,
	from time.Time,
	to time.Time,
) (ytrackercore.WorkLogs, error) {
	const method = "/worklogs"

	handleError := httpcore.GetHandleErrorFunc[ytrackercore.WorkLogs](serviceName, method, nil)

	req := adapter.client.NewRequest("GET", method, nil).
		SetQueryParamsFromValues(port.GetGetWorkLogsRequest(createdBy, from, to))

	var respData port.WorkLogsResponse

	resp, err := adapter.client.Do(req, &respData)
	if err != nil {
		return handleError(err, resp)
	}

	if err := httpcore.HandleHTTPError(err, resp); err != nil {
		return handleError(err, resp)
	}

	result, err := respData.ToCore()
	if err != nil {
		return handleError(err, resp)
	}

	return result, nil
}

func (adapter *Adapter) AddWorkLogs(
	_ context.Context,
	taskKey string,
	start time.Time,
	duration time.Duration,
	comment string,
) (ytrackercore.WorkLog, error) {
	const method = "/v2/issues/{taskKey}/worklog"

	handleError := httpcore.GetHandleErrorFunc[ytrackercore.WorkLog](serviceName, method, ytrackercore.WorkLog{})

	pathParams, body := port.GetAddWorkLogsRequest(taskKey, start, duration, comment)

	req := adapter.client.NewRequest("POST", method, nil).
		SetPathParams(pathParams).
		SetBody(body)

	var respData port.WorkLogResponse

	resp, err := adapter.client.Do(req, &respData)
	if err != nil {
		return handleError(err, resp)
	}

	if err := httpcore.HandleHTTPError(err, resp); err != nil {
		return handleError(err, resp)
	}

	result, err := respData.ToCore()
	if err != nil {
		return handleError(err, resp)
	}

	return result, nil
}
