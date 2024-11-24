package usecase

import (
	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/service"
	"github.com/emildeev/harvest-yt/internal/usecase/migrator"
	"github.com/emildeev/harvest-yt/internal/usecase/validator"
)

type Provider struct {
	Migrator  *migrator.UseCase
	Validator *validator.UseCase
}

func New(_ config.Config, provider *service.Provider) (*Provider, error) {
	return &Provider{
		Migrator:  migrator.New(provider.YTracker, provider.Harvest, provider.TimeTable),
		Validator: validator.New(provider.YTracker, provider.Harvest),
	}, nil
}
