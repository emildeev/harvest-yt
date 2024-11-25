package migrator

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	harvestcore "github.com/emildeev/harvest-yt/internal/core/harvest"
	timetable "github.com/emildeev/harvest-yt/internal/core/time_table"
)

func (uc *UseCase) GetList(ctx context.Context) (timetable.Table, error) {
	pushed, notPushed, err := uc.harvestService.GetTimersList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get timers list: %w", err)
	}

	offset := getOffset(pushed)

	taskTable := uc.timeTable.Generate(ctx, notPushed, offset)

	return taskTable, nil
}

func getOffset(pushed harvestcore.TimeEntries) (offset time.Duration) {
	for _, timer := range pushed {
		offset += timer.Hours
	}
	return offset
}

func (uc *UseCase) SpendTime(ctx context.Context, table timetable.Table) error {
	for _, element := range table {
		if element.Err != nil {
			continue
		}
		logID, err := uc.yTrackerService.SpendTime(
			ctx, element.TaskKey, element.StartTime, element.Duration, element.Comment,
		)
		if err != nil {
			slog.Error(fmt.Errorf("failed to spend time: %w", err).Error())
			continue
		}

		err = uc.harvestService.MarkTimerPushed(ctx, element.TimerID, logID)
		if err != nil {
			slog.Error(fmt.Errorf("failed to mark timer pushed: %w", err).Error())
		}
	}
	return nil
}
