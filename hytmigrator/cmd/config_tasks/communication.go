package configtaskscmd

import (
	"context"
	"fmt"
	"maps"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"

	"github.com/emildeev/harvest-yt/hytmigrator/cmd/helper"
	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

func HandleGetCommunicationTasks(
	_ context.Context,
	currentCfg config.Config,
	_ *usecase.Provider,
) (config.Config, error) {
	tableRowHeader := table.Row{"Harvest Task", "Task Key"}
	communicationTasks := currentCfg.Tasks.CommunicationTasks

	tableRows := make([]table.Row, len(communicationTasks))
	i := 0
	for harvestTask, TaskKey := range communicationTasks {
		tableRows[i] = table.Row{
			harvestTask,
			TaskKey,
		}
		i++
	}

	tw := table.NewWriter()
	tw.AppendHeader(tableRowHeader)
	tw.AppendRows(tableRows)
	fmt.Println(tw.Render())
	return currentCfg, nil
}

func HandleUpdateCommunicationTasks(
	ctx context.Context,
	currentCfg config.Config,
	provider *usecase.Provider,
) (config.Config, error) {
	newConfig := currentCfg
	newConfig.Tasks.CommunicationTasks = maps.Clone(currentCfg.Tasks.CommunicationTasks)
	_, _ = HandleGetCommunicationTasks(ctx, newConfig, provider)

	newCommunicationTasks := newConfig.Tasks.CommunicationTasks

	if newCommunicationTasks == nil {
		newCommunicationTasks = make(map[string]string)
	}

	harvestTaskPromptValidation := func(input string) error {
		return provider.Validator.ValidateHarvestTask(ctx, input)
	}

	harvestTaskPrompt := promptui.Prompt{
		Label:    "Harvest task for edit or new",
		Validate: harvestTaskPromptValidation,
	}

	harvestTaskToEdit, err := harvestTaskPrompt.Run()
	if err != nil {
		return currentCfg, err
	}

	harvestTaskToEdit = strings.ToLower(harvestTaskToEdit)

	taskKey, err := getTrackerTaskKey(ctx, newCommunicationTasks[harvestTaskToEdit], provider)
	if err != nil {
		return currentCfg, err
	}

	if taskKey != "" {
		newCommunicationTasks[harvestTaskToEdit] = taskKey
	} else {
		if _, ok := newCommunicationTasks[harvestTaskToEdit]; ok {
			delete(newCommunicationTasks, harvestTaskToEdit)
		}
	}

	newConfig.Tasks.CommunicationTasks = newCommunicationTasks

	_, _ = HandleGetCommunicationTasks(ctx, newConfig, provider)

	confirm := helper.GetConfirmation()
	if !confirm {
		fmt.Println("Aborted")
		return currentCfg, nil
	}

	fmt.Println("Save config")

	return newConfig, nil
}
