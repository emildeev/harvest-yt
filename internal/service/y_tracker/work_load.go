package ytracker

import (
	"context"
	"fmt"
	"time"

	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func (service *Service) GetTodayTimeTable(ctx context.Context) (ytrackercore.WorkLogs, error) {
	user, err := service.adapter.GetMyself(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get myself: %w", err)
	}

	fromTime := time.Now().Truncate(24 * time.Hour)
	toTime := fromTime.Add(24 * time.Hour)

	workLogs, err := service.adapter.GetWorkLogs(ctx, user.Login, fromTime, toTime)
	if err != nil {
		return nil, fmt.Errorf("failed to get work logs: %w", err)
	}
	return workLogs, nil
}

func (service *Service) SpendTime(
	ctx context.Context,
	taskKey string,
	startTime time.Time,
	duration time.Duration,
	comment string,
) (int, error) {
	log, err := service.adapter.AddWorkLogs(ctx, taskKey, startTime, duration, comment)
	if err != nil {
		return 0, fmt.Errorf("failed to add work log: %w", err)
	}
	return log.ID, nil
}
