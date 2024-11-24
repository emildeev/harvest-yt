package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/emildeev/harvest-yt/pkg/helper"
)

const (
	DefaultLogLevel = slog.LevelError
	LogLevelLen     = 3
)

var (
	rootCmd = &cobra.Command{
		Use:              "hytmigrator",
		Short:            "Harvest Yandex Tracker Migrator",
		Long:             `This tool is used for migrating timers from Harvest to Yandex Tracker time spend.`,
		TraverseChildren: true,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	logLevelUsage := fmt.Sprintf(
		"set app log level (%s, %s, %s, %s) default is off",
		slog.LevelError.String(), slog.LevelWarn.String(), slog.LevelInfo.String(), slog.LevelDebug.String(),
	)

	rootCmd.PersistentFlags().StringP(
		"log_level", "l", "", logLevelUsage,
	)
}

func initConfig() {
	var err error
	time.Local, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".hytmigrator")
	_ = viper.ReadInConfig()

	InitLogger()
}

func InitLogger() {
	logLevel, _ := rootCmd.PersistentFlags().GetString("log_level")
	slog.SetLogLoggerLevel(getSlogLogLevel(logLevel))
	if logLevel == "" {
		logger := slog.New(slog.NewTextHandler(io.Discard, nil))
		slog.SetDefault(logger)
	}
}

func getSlogLogLevel(strLevel string) slog.Level {
	strLevel = makeLogLevelStr(strLevel)

	switch strLevel {
	case makeLogLevelStr(slog.LevelError.String()):
		return slog.LevelError
	case makeLogLevelStr(slog.LevelWarn.String()):
		return slog.LevelWarn
	case makeLogLevelStr(slog.LevelInfo.String()):
		return slog.LevelInfo
	case makeLogLevelStr(slog.LevelDebug.String()):
		return slog.LevelDebug
	default:
		return DefaultLogLevel
	}
}

func makeLogLevelStr(s string) string {
	return strings.ToLower(helper.StringTruncate(s, LogLevelLen))
}
