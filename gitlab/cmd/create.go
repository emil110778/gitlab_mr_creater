/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/emildeev/gitlab_helper/internal"
	"github.com/emildeev/gitlab_helper/internal/config"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/spf13/cobra"
)

const (
	additionalBrunchFlag = "additional_brunch"
	mainBrunchFlag       = "main_brunch"
)

// creteCmd represents the crete command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create merge requests",
	Long: `This command get required information from git and yandex tracker and ask addition information in cli for creating
merge requests from current brunch to target main brunch (configured) and additional brunches (configured).
After creating merge requests it will show links to created merge requests and will set it to ticket in yandex tracker`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		cfg, err := config.New()
		if err != nil {
			slog.Error("error load config err", err.Error())
			return ErrConfigure
		}

		provider, err := internal.New(cfg)
		if err != nil {
			slog.Error("error configure provider err", err.Error())
			return InternalErr
		}

		createAdditional, _ := cmd.Flags().GetBool(additionalBrunchFlag)

		mrs, err := provider.MR.Create(ctx, createAdditional)
		if err != nil {
			slog.Error("error create MR err", err.Error())
			return InternalErr
		}

		logMRs(mrs)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	const additionalBrunchUsage = "create merge requests to additional brunches"
	createCmd.Flags().BoolP(
		additionalBrunchFlag, "a", false, additionalBrunchUsage,
	)
}

func logMRs(mrs []gitlabcore.ResultMRInfo) {
	for _, mr := range mrs {
		log := fmt.Sprintf("\nMR to brunch: %s", mr.Brunch)
		if mr.URL != "" {
			log += fmt.Sprintf("\nurl: %s", mr.URL)
		}
		if mr.Err != nil {
			log += fmt.Sprintf("\nerror: %s", mr.Err.Error())
		}
		fmt.Println(log)
	}
}
