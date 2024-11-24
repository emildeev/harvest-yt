package configtaskscmd

import (
	"context"

	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

func HandleGetSkippedTasks(_ context.Context, currentCfg config.Config, _ *usecase.Provider) (config.Config, error) {
	handleGetTaskList(currentCfg.Tasks.SkippedTasks)
	return currentCfg, nil
}

func HandleUpdateSkippedTasks(
	ctx context.Context,
	currentCfg config.Config,
	provider *usecase.Provider,
) (config.Config, error) {
	var err error
	currentCfg.Tasks.SkippedTasks, err = handleUpdateTaskList(ctx, currentCfg.Tasks.SkippedTasks, provider)
	return currentCfg, err
}
