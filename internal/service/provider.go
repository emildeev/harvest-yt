package service

import (
	"github.com/emildeev/harvest-yt/internal/adapter"
	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/service/harvest"
	timetable "github.com/emildeev/harvest-yt/internal/service/time_table"
	ytracker "github.com/emildeev/harvest-yt/internal/service/y_tracker"
)

type Provider struct {
	YTracker  *ytracker.Service
	Harvest   *harvest.Service
	TimeTable *timetable.Service
}

func New(cfg config.Config, provider adapter.Provider) (*Provider, error) {
	yTracker := ytracker.New(provider.HTTP.YTracker)

	return &Provider{
		YTracker:  yTracker,
		Harvest:   harvest.New(provider.HTTP.Harvest),
		TimeTable: timetable.New(yTracker, cfg.Tasks),
	}, nil
}
