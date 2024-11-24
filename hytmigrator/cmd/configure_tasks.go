/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	configtaskscmd "github.com/emildeev/harvest-yt/hytmigrator/cmd/config_tasks"
	"github.com/emildeev/harvest-yt/internal"
	"github.com/emildeev/harvest-yt/internal/config"
	"github.com/emildeev/harvest-yt/internal/usecase"
)

const (
	getSkippedTasksArg          = "skipped"
	updateSkippedTasksArg       = "skipped_update"
	getDeveloperTasksArg        = "developer"
	updateDeveloperTasksArg     = "developer_update"
	getCommunicationTasksArg    = "communication"
	updateCommunicationTasksArg = "communication_update"
)

// configureCmd represents the configure command
var configureTasksCmd = &cobra.Command{
	Use:   "configure_tasks",
	Short: "Configure tasks for tool",
	Long: `This command will configure tasks for tool:
develop harvest tasks and communication harvest tasks to tracker tasks mapping
`,
	Args: cobra.ExactArgs(1),
	ValidArgs: []string{
		getSkippedTasksArg,
		updateSkippedTasksArg,
		getDeveloperTasksArg,
		updateDeveloperTasksArg,
		getCommunicationTasksArg,
		updateCommunicationTasksArg,
	},
	SilenceErrors: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cobra.OnlyValidArgs(cmd, args)
		if err != nil {
			return err
		}

		ctx := cmd.Context()

		currentCfg := config.NewWithoutValidate()

		handlers := map[string]func(
			ctx context.Context,
			currentCfg config.Config,
			provider *usecase.Provider,
		) (config.Config, error){
			getSkippedTasksArg:          configtaskscmd.HandleGetSkippedTasks,
			updateSkippedTasksArg:       configtaskscmd.HandleUpdateSkippedTasks,
			getDeveloperTasksArg:        configtaskscmd.HandleGetDeveloperTasks,
			updateDeveloperTasksArg:     configtaskscmd.HandleUpdateDeveloperTasks,
			getCommunicationTasksArg:    configtaskscmd.HandleGetCommunicationTasks,
			updateCommunicationTasksArg: configtaskscmd.HandleUpdateCommunicationTasks,
		}

		provider, err := internal.New(ctx, currentCfg)
		if err != nil {
			slog.Error("configure provider:", "err", err.Error())
			return InternalErr
		}

		cfg, err := handlers[args[0]](ctx, currentCfg, provider)
		if err != nil {
			return err
		}

		configMap := make(map[string]interface{})
		err = mapstructure.Decode(cfg, &configMap)
		if err != nil {
			return err
		}

		for key, val := range configMap {
			viper.Set(key, val)
		}

		err = viper.WriteConfig()
		if err != nil {
			if errors.As(err, &viper.ConfigFileNotFoundError{}) {
				err = viper.SafeWriteConfig()
				if err != nil {
					return fmt.Errorf("error crete file: %w", err)
				}
			} else {
				return fmt.Errorf("error save file: %w", err)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureTasksCmd)
}
