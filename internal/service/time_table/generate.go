package timetable

import (
	"context"
	"fmt"
	"strings"
	"time"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
	"github.com/emildeev/harvest-yt/internal/core/time_table"
	"github.com/emildeev/harvest-yt/pkg/helper"
)

func (service *Service) Generate(ctx context.Context, entries harvestcore.TimeEntries) timetablecore.Table {
	developerSlice := helper.GetMapFromSlice(service.cfg.DevelopTasks)

	result := make(timetablecore.Table, 0, len(entries))

	currentStartTime, _ := time.Parse(time.TimeOnly, service.cfg.StartTime)
	dateNow := time.Now().Truncate(24 * time.Hour).Add(-24 * time.Hour)
	currentStartTime = currentStartTime.AddDate(dateNow.Year(), int(dateNow.Month()), dateNow.Day())
	for _, entry := range entries {
		var task timetablecore.Element
		if _, ok := developerSlice[entry.Task]; ok {
			task = service.getForDeveloperTask(entry)
		} else {
			task = service.getForCommunication(entry)
		}

		taskInfo, err := service.yTrackerService.ValidateTicketForSpend(ctx, task.TaskKey)
		if err != nil {
			if task.Err == nil {
				task.Err = fmt.Errorf("GetTicket:%w", err)
			}
		}
		task.TaskTitle = taskInfo.Title

		task.StartTime = currentStartTime
		result = append(result, task)

		currentStartTime = currentStartTime.Add(task.Duration)
	}
	return result
}

func (service *Service) getForDeveloperTask(entry harvestcore.TimeEntry) timetablecore.Element {
	element := timetablecore.Element{
		Duration: entry.Hours,
	}

	taskKey := service.taskKeyRegexp.Find([]byte(entry.Notes))
	if len(taskKey) != 0 {
		element.TaskKey = strings.Replace(string(taskKey), ":", "", -1)
	} else {
		element.Err = fmt.Errorf("task key not found in notes")
		element.Comment = entry.Notes
	}

	return element
}

func (service *Service) getForCommunication(entry harvestcore.TimeEntry) timetablecore.Element {
	element := timetablecore.Element{
		Comment:  entry.Notes,
		Duration: entry.Hours,
	}

	if taskKey, ok := service.cfg.CommunicationTasks[entry.Task]; ok {
		element.TaskKey = taskKey
	} else {
		element.Err = fmt.Errorf("task key not found in communication")
	}

	return element
}
