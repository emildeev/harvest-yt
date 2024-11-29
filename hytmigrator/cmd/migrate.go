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

		confirm := helper.GetConfirmation()
		if !confirm {
			fmt.Println("Migration canceled")
			return nil
		}

		err = provider.Migrator.SpendTime(ctx, taskTable)
		if err != nil {
			return err
		}

		fmt.Println("Migration completed")

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
		taskTitle := task.TaskTitle

		if len([]rune(taskTitle)) > 50 {
			taskTitle = string([]rune(taskTitle)[:50]) + "..."
		}
		comment := task.Comment
		if len([]rune(comment)) > 40 {
			comment = string([]rune(comment)[:40]) + "..."
		}

		tableRows[i] = table.Row{
			task.TaskKey,
			taskTitle,
			comment,
			task.Duration,
			task.StartTime.Format("2006-01-02 15:04:05"),
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
