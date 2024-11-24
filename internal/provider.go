package internal

import (
	"context"

	"github.com/emildeev/harvest-yt/internal/adapter"
	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/connection"
	"github.com/emildeev/harvest-yt/internal/service"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

func New(ctx context.Context, cfg config.Config) (provider *usecase.Provider, err error) {
	connectionProvider, err := connection.New(ctx, cfg)
	if err != nil {
		return provider, err
	}
	drivenProvider, err := adapter.New(connectionProvider)
	if err != nil {
		return provider, err
	}
	serviceProvider, err := service.New(cfg, drivenProvider)
	if err != nil {
		return provider, err
	}
	useCaseProvider, err := usecase.New(cfg, serviceProvider)
	if err != nil {
		return provider, err
	}
	return useCaseProvider, nil
}
