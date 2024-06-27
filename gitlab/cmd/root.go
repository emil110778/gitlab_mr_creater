/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/emildeev/gitlab_helper/pkg/helper"
)

const (
	DefaultLogLevel = slog.LevelError
	LogLevelLen     = 3
)

var (
	rootCmd = &cobra.Command{
		Use:   "gitlab",
		Short: "Gitlab helper tool",
		Long: `This cli application is used to create merge requests in gitlab
It uses gitlab api, yandex tracker api to create merge request with Title and Description,
and git commands for getting current branch and repository`,
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
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".fullstack")
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
