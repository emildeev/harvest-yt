package timetable

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
	"github.com/emildeev/harvest-yt/internal/core/time_table"
	"github.com/emildeev/harvest-yt/pkg/helper"
)

func (service *Service) Generate(
	ctx context.Context,
	entries harvestcore.TimeEntries,
	offset time.Duration,
) timetablecore.Table {
	developerSlice := helper.GetMapFromSlice(service.cfg.DevelopTasks)

	result := make(timetablecore.Table, 0, len(entries))

	dateNow := carbon.Now()
	startTime := carbon.ParseByLayout(service.cfg.StartTime, time.TimeOnly, dateNow.Location())
	startTime = startTime.SetDate(dateNow.Year(), dateNow.Month(), dateNow.Day())
	currentStartTime := startTime.StdTime()

	for _, entry := range entries {
		task := timetablecore.Element{}
		if _, ok := developerSlice[entry.Task]; ok {
			task = service.getForDeveloperTask(entry)
		} else {
			task = service.getForCommunication(entry)
		}

		task.TimerID = entry.ID

		taskInfo, err := service.yTrackerService.ValidateTicketForSpend(ctx, task.TaskKey)
		if err != nil {
			if task.Err == nil {
				task.Err = fmt.Errorf("GetTicket:%w", err)
			}
		}
		task.TaskTitle = taskInfo.Title

		task.StartTime = currentStartTime
		if task.Err == nil {
			currentStartTime = currentStartTime.Add(task.Duration)
		}

		result = append(result, task)
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
