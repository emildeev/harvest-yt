package harvest

import (
	"context"

	"github.com/emildeev/harvest-yt/pkg/helper"
)

func (service *Service) GetTasksListMap(ctx context.Context) (map[string]struct{}, error) {
	if service.tasksCache != nil {
		return helper.GetMapFromSlice(service.tasksCache), nil
	}

	tasks, err := service.adapter.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	service.tasksCache = tasks

	return helper.GetMapFromSlice(tasks), nil
}
