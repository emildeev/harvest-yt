package harvest

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/emildeev/harvest-yt/internal/adapter/http/harvest/port"
	httpcore "github.com/emildeev/harvest-yt/internal/core/http"
)

func (adapter *Adapter) GetAllTasks(ctx context.Context) ([]string, error) {
	handleError := httpcore.GetHandleErrorFunc[[]string](
		serviceName, "GetTasksList", nil,
	)

	tasks, resp, err := adapter.tasks.List(ctx, port.GetTasksListRequest())
	response := &resty.Response{
		RawResponse: resp,
	}
	if err != nil {
		return handleError(err, response)
	}

	if err := httpcore.HandleHTTPError(err, response); err != nil {
		return handleError(err, response)
	}

	return port.GetTaskListResponse(tasks), nil
}
