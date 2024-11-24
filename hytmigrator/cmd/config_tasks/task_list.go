package configtaskscmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/manifoldco/promptui"
	"golang.org/x/exp/slices"

	"github.com/emildeev/harvest-yt/hytmigrator/cmd/helper"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

func handleGetTaskList(currentCfg []string) {
	tableRowHeader := table.Row{"ID", "Harvest Task"}

	tableRows := make([]table.Row, len(currentCfg))
	for i, task := range currentCfg {
		tableRows[i] = table.Row{
			i,
			task,
		}
	}

	tw := table.NewWriter()
	tw.AppendHeader(tableRowHeader)
	tw.AppendRows(tableRows)
	fmt.Println(tw.Render())
}

func handleUpdateTaskList(ctx context.Context, currentCfg []string, provider *usecase.Provider) ([]string, error) {
	newConfig := slices.Clone(currentCfg)
	newConfig = append(newConfig, "(new task)")
	handleGetTaskList(newConfig)

	validateNumberPrompt := func(input string) error {
		num, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("it should be number number")
		}
		if num > len(newConfig)-1 {
			return errors.New("it should be number from list")
		}
		return nil
	}

	numberPrompt := promptui.Prompt{
		Label:    "ID for edit",
		Validate: validateNumberPrompt,
	}

	idStr, err := numberPrompt.Run()
	if err != nil {
		return currentCfg, err
	}

	id, _ := strconv.Atoi(idStr)

	if id != len(newConfig)-1 {
		newConfig = newConfig[:len(newConfig)-1]
	} else {
		newConfig[len(newConfig)-1] = ""
	}

	harvestTaskPromptValidate := func(input string) error {
		if input == "" {
			return nil
		}
		return provider.Validator.ValidateHarvestTask(ctx, input)
	}

	harvestTaskPrompt := promptui.Prompt{
		Label:    "Harvest task (empty for delete)",
		Default:  newConfig[id],
		Validate: harvestTaskPromptValidate,
	}

	harvestTask, err := harvestTaskPrompt.Run()
	if err != nil {
		return currentCfg, err
	}

	if harvestTask != "" {
		newConfig[id] = strings.ToLower(harvestTask)
	} else {
		newConfig = append(newConfig[:id], newConfig[id+1:]...)
	}

	handleGetTaskList(newConfig)

	confirm := helper.GetConfirmation()
	if !confirm {
		fmt.Println("Aborted")
		return currentCfg, nil
	}

	fmt.Println("Save config")

	return newConfig, nil
}
