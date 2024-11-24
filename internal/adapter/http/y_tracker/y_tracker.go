package ytracker

import (
	tracker "github.com/emildeev/yandex-tracker-go"
	"github.com/go-resty/resty/v2"
)

const (
	serviceName = "y_tracker"
)

type Client interface {
	GetTicket(ticketKey string) (ticket tracker.Ticket, err error)
	PatchTicket(ticketKey string, body map[string]string) (ticket tracker.Ticket, err error)
	GetTicketComments(ticketKey string) (comments tracker.TicketComments, err error)
	Myself() (user *tracker.User, err error)
	NewRequest(method, path string, opt any) *resty.Request
	Do(req *resty.Request, v any) (*resty.Response, error)
}

type Adapter struct {
	client Client
}

func New(client Client) *Adapter {
	return &Adapter{
		client: client,
	}
}
