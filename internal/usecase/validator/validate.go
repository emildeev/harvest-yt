package validator

import (
	"context"
	"fmt"
	"strings"
)

func (uc *UseCase) ValidateHarvestTask(ctx context.Context, taskName string) error {
	taskName = strings.ToLower(taskName)

	tasks, err := uc.harvestService.GetTasksListMap(ctx)
	if err != nil {
		return fmt.Errorf("failed to get tasks list: %w", err)
	}

	if _, ok := tasks[taskName]; !ok {
		return fmt.Errorf("task %s not found", taskName)
	}

	return nil
}

func (uc *UseCase) ValidateYTrackerTicket(ctx context.Context, ticketKey string) error {
	_, err := uc.yTrackerService.GetTicketTitle(ctx, ticketKey)
	if err != nil {
		return fmt.Errorf("failed to get ticket title: %w", err)
	}

	return nil
}
