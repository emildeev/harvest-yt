package port

import (
	"strings"
	"time"

	"github.com/emildeev/go-harvest/harvest"
	"github.com/golang-module/carbon/v2"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
)

func GetTodayOpt() *harvest.TimeEntryListOptions {
	return &harvest.TimeEntryListOptions{
		From: &harvest.Date{Time: carbon.Now().StartOfDay().StdTime()},
		To:   &harvest.Date{Time: carbon.Now().EndOfDay().StdTime()},
	}
}

func GetTimeEntriesFormTimeEntryList(list *harvest.TimeEntryList) harvestcore.TimeEntries {
	if list == nil || list.TimeEntries == nil {
		return nil
	}
	timeEntries := make(harvestcore.TimeEntries, 0, len(list.TimeEntries))
	for _, entry := range list.TimeEntries {
		if entry != nil {
			timeEntries = append(timeEntries, GetTimeEntryFromTimeEntry(entry))
		}
	}
	return timeEntries
}

func GetTimeEntryFromTimeEntry(entry *harvest.TimeEntry) harvestcore.TimeEntry {
	res := harvestcore.TimeEntry{}
	if entry.ID != nil {
		res.ID = *entry.ID
	}
	if entry.Client != nil && entry.Client.Name != nil {
		res.Client = *entry.Client.Name
	}
	if entry.Project != nil && entry.Project.Name != nil {
		res.Project = *entry.Project.Name
	}
	if entry.Task != nil && entry.Task.Name != nil {
		res.Task = strings.ToLower(*entry.Task.Name)
	}
	if entry.Hours != nil {
		res.Hours = time.Duration(*entry.Hours * float64(time.Hour))
	}
	if entry.RoundedHours != nil {
		res.RoundedHours = time.Duration(*entry.RoundedHours * float64(time.Hour))
	}
	if entry.Notes != nil {
		res.Notes = *entry.Notes
	}
	if entry.IsRunning != nil {
		res.IsRunning = *entry.IsRunning
	}
	return res
}

func GetTimeEntryUpdateCommentRequest(comment string) *harvest.TimeEntryUpdate {
	return &harvest.TimeEntryUpdate{
		Notes: &comment,
	}
}
