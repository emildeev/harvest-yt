package ytracker

import (
	"context"

	"github.com/emildeev/harvest-yt/internal/adapter/http/y_tracker/port"
	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func (adapter *Adapter) GetTicket(_ context.Context, ticketKey string) (ticket ytrackercore.Ticket, err error) {
	rawTicket, err := adapter.client.GetTicket(ticketKey)
	if err != nil {
		return ticket, err
	}
	return port.GetGetTicketResponse(rawTicket), nil
}
