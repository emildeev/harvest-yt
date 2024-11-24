/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"

	"github.com/emildeev/harvest-yt/hytmigrator/cmd/helper"
	"github.com/emildeev/harvest-yt/internal"
	"github.com/emildeev/harvest-yt/internal/config"
	timetable "github.com/emildeev/harvest-yt/internal/core/time_table"
)

// creteCmd represents the crete command
var createCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate harvest timers to yandex tracker",
	Long: `This command migrate hearvest timers to yandex tracker spend times, 
it use configured developer task for get task for get yandex tracker task key 
or communication tasks for get yandex tracker task key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		cfg, err := config.New()
		if err != nil {
			slog.Error("load config:", "err", err.Error())
			return ErrConfigure
		}

		provider, err := internal.New(ctx, cfg)
		if err != nil {
			slog.Error("configure provider:", "err", err.Error())
			return InternalErr
		}

		taskTable, err := provider.Migrator.GetList(ctx)
		if err != nil {
			slog.Error("migrate", "err", err.Error())
			return InternalErr
		}

		printTaskTable(taskTable)

		helper.GetConfirmation()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func printTaskTable(taskTable timetable.Table) {
	tableRowHeader := table.Row{"Task Key", "Task Title", "Comment", "Duration", "Start Time"}
	hasErr := taskTable.HasErr()
	if hasErr {
		tableRowHeader = append(tableRowHeader, "Error")
	}

	tableRows := make([]table.Row, len(taskTable))
	totalTime := time.Duration(0)
	for i, task := range taskTable {
		tableRows[i] = table.Row{
			task.TaskKey,
			task.TaskTitle,
			task.Comment,
			task.Duration,
			task.StartTime,
		}
		if hasErr {
			tableRows[i] = append(tableRows[i], task.Err)
		}
		totalTime += task.Duration
	}
	tableRowFooter := table.Row{"", "", "Total Time", totalTime}

	tw := table.NewWriter()
	tw.AppendHeader(tableRowHeader)
	tw.AppendRows(tableRows)
	tw.AppendFooter(tableRowFooter)
	fmt.Println(tw.Render())
}
