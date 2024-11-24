package http

import (
	"errors"
	"fmt"

	"github.com/emildeev/harvest-yt/internal/adapter/http/harvest"
	ytracker "github.com/emildeev/harvest-yt/internal/adapter/http/y_tracker"
	httpconn "github.com/emildeev/harvest-yt/internal/connection/http"
)

type Provider struct {
	YTracker *ytracker.Adapter
	Harvest  *harvest.Adapter
}

func New(connections *httpconn.Connection) (*Provider, error) {
	errorWrapper := func(err error) (*Provider, error) {
		return nil, fmt.Errorf("error configure http adapter %w", err)
	}

	if connections == nil {
		return errorWrapper(errors.New("http connections is nil"))
	}

	yTrackerAdapter := ytracker.New(connections.YTracker)

	harvestAdapter := harvest.New(connections.Harvest.Timesheet, connections.Harvest.Task)
	return &Provider{
		YTracker: yTrackerAdapter,
		Harvest:  harvestAdapter,
	}, nil
}
