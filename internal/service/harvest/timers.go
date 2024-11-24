package harvest

import (
	"context"
	"fmt"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
)

func (service *Service) GetTimersList(
	ctx context.Context,
) (pushed harvestcore.TimeEntries, notPushed harvestcore.TimeEntries, err error) {
	timers, err := service.adapter.GetTimesheetListToday(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get timers: %w", err)
	}
	pushed, notPushed = service.filterSplitPushed(timers)

	return pushed, notPushed, nil
}

func (service *Service) filterSplitPushed(
	timers harvestcore.TimeEntries,
) (pushed harvestcore.TimeEntries, notPushed harvestcore.TimeEntries) {
	for _, timer := range timers {
		if service.pushedRegexp.MatchString(timer.Notes) {
			pushed = append(pushed, timer)
		} else {
			notPushed = append(notPushed, timer)
		}
	}
	return pushed, notPushed
}

func (service *Service) MarkTimerPushed(ctx context.Context, id int64, workloadId int) error {
	timer, err := service.adapter.GetTimesheet(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get timer: %w", err)
	}
	_, err = service.adapter.UpdateTimesheetComment(ctx, id, fmt.Sprintf(pushFormat, workloadId)+timer.Notes)
	if err != nil {
		return fmt.Errorf("failed to update timer: %w", err)
	}
	return nil
}
