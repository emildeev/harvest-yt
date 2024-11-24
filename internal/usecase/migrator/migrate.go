package migrator

import (
	"context"
	"fmt"

	timetable "github.com/emildeev/harvest-yt/internal/core/time_table"
)

func (uc *UseCase) GetList(ctx context.Context) (timetable.Table, error) {
	_, notPushed, err := uc.harvestService.GetTimersList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get timers list: %w", err)
	}

	taskTable := uc.timeTable.Generate(ctx, notPushed)

	return taskTable, nil
}
