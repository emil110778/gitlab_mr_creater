/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/emil110778/gitlab_mr_creator/internal"
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
	"github.com/spf13/cobra"
)

// creteCmd represents the crete command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create merge requests",
	Long:  `This command create merge request to prod and release (if set flag -release)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		cfg, err := config.New()
		if err != nil {
			slog.Debug("error load config err", err.Error())
			return ErrConfigure
		}

		provider, err := internal.New(cfg)
		if err != nil {
			slog.Debug("error configure provider err", err.Error())
			return InternalErr
		}

		mrs, err := provider.MR.Create(ctx)
		if err != nil {
			slog.Debug("error create MR err", err.Error())
			return InternalErr
		}

		logMRs(mrs)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// creteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// creteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logMRs(mrs []gitlabcore.CreatedMRInfo) {
	for _, mr := range mrs {
		log := fmt.Sprintf("MR to brunch: %s", mr.Brunch)
		if mr.URL != "" {
			log += fmt.Sprintf("\nurl: %s", mr.URL)
		}
		if mr.Err != nil {
			log += fmt.Sprintf("\nerror: %s", mr.Err.Error())
		}
		fmt.Println(log)
	}
}
