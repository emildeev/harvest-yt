package ytracker

import (
	"context"
	"errors"

	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func (service *Service) GetTicket(ctx context.Context, ticketKey string) (ticket ytrackercore.Ticket, err error) {
	return service.adapter.GetTicket(ctx, ticketKey)
}

func (service *Service) GetTicketTitle(ctx context.Context, ticketKey string) (string, error) {
	ticket, err := service.GetTicket(ctx, ticketKey)
	if err != nil {
		return "", err
	}

	return ticket.Title, nil
}

func (service *Service) ValidateTicketForSpend(ctx context.Context, taskKey string) (ytrackercore.Ticket, error) {
	ticket, err := service.GetTicket(ctx, taskKey)
	if err != nil {
		return ytrackercore.Ticket{}, err
	}

	err = validateTaskForSpendTime(ctx, ticket)
	if err != nil {
		return ticket, err
	}

	return ticket, nil
}

func validateTaskForSpendTime(_ context.Context, task ytrackercore.Ticket) error {
	if task.Customer == "" {
		return errors.New("customer is empty")
	}
	switch task.Status {
	case ytrackercore.DoneStatus:
		return errors.New("task is done")
	case ytrackercore.RejectedStatus:
		return errors.New("task is rejected")
	case ytrackercore.PauseStatus:
		return errors.New("task is paused")
	}
	return nil
}
