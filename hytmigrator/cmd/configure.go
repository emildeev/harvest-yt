/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/becoded/go-harvest/harvest"
	tracker "github.com/emildeev/yandex-tracker-go"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	configcmd "github.com/emildeev/harvest-yt/hytmigrator/cmd/config"
	"github.com/emildeev/harvest-yt/internal/config"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure environment for tool",
	Long: `This command will configure environment for tool:
hytmigrator credentials
yandex tracker credentials
and repository branch configuration
`,
	SilenceErrors: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		currentCfg := config.NewWithoutValidate()

		fmt.Println()
		fmt.Println(
			"for generate Harvest API token and get account ID:\n" +
				"https://id.getharvest.com/oauth2/access_tokens/new",
		)
		fmt.Println()
		accountID, err := configcmd.HarvestAccountID(currentCfg.HTTP.Harvest.AccountID)
		if err != nil {
			return err
		}

		token, err := configcmd.HarvestGetToken(currentCfg.HTTP.Harvest.Token)
		if err != nil {
			return err
		}

		client := harvest.NewAPIClient(
			oauth2.NewClient(
				ctx, oauth2.StaticTokenSource(
					&oauth2.Token{
						AccessToken: token,
					},
				),
			),
		)
		client.AccountID = accountID

		_, resp, err := client.User.Current(ctx)
		if err != nil {
			if resp != nil && resp.StatusCode != http.StatusOK {
				slog.Error("hytmigrator authorization:", "err", err)
				return errors.New("hytmigrator authorization error")
			}
			return err
		}

		yTrackerOrgID, err := configcmd.YTrackerGetOrgID(currentCfg.HTTP.YTracker.OrgID)
		if err != nil {
			return err
		}

		yTrackerToken, err := configcmd.YTrackerGetToken(currentCfg.HTTP.YTracker.Token)
		if err != nil {
			return err
		}

		yTrackerClient := tracker.New("OAuth "+yTrackerToken, strconv.Itoa(yTrackerOrgID), "")
		_, err = yTrackerClient.Myself()
		if err != nil {
			return err
		}

		startTime, err := configcmd.GetStartTime(currentCfg.Tasks.StartTime)
		if err != nil {
			return err
		}

		cfg := currentCfg

		cfg.HTTP.Harvest.AccountID = accountID
		cfg.HTTP.Harvest.Token = token

		cfg.HTTP.YTracker.Token = yTrackerToken
		cfg.HTTP.YTracker.OrgID = yTrackerOrgID

		cfg.Tasks.StartTime = startTime
		fmt.Println(cfg.Tasks.StartTime)

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
	rootCmd.AddCommand(configureCmd)
}
