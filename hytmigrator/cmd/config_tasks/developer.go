package configtaskscmd

import (
	"context"

	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

func HandleGetDeveloperTasks(_ context.Context, currentCfg config.Config, _ *usecase.Provider) (config.Config, error) {
	handleGetTaskList(currentCfg.Tasks.DevelopTasks)
	return currentCfg, nil
}

func HandleUpdateDeveloperTasks(
	ctx context.Context,
	currentCfg config.Config,
	provider *usecase.Provider,
) (config.Config, error) {
	var err error
	currentCfg.Tasks.DevelopTasks, err = handleUpdateTaskList(ctx, currentCfg.Tasks.DevelopTasks, provider)
	return currentCfg, err
}
