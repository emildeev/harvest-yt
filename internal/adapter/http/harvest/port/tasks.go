package port

import (
	"strings"

	"github.com/becoded/go-harvest/harvest"
)

func GetTasksListRequest() *harvest.TaskListOptions {
	return &harvest.TaskListOptions{
		IsActive: true,
	}
}

func GetTaskListResponse(resp *harvest.TaskList) []string {
	if resp == nil || resp.Tasks == nil {
		return nil
	}
	res := make([]string, 0, len(resp.Tasks))
	for _, task := range resp.Tasks {
		if task != nil && task.Name != nil {
			res = append(res, strings.ToLower(*task.Name))
		}
	}
	return res
}
