/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/emildeev/gitlab_helper/internal"
	"github.com/emildeev/gitlab_helper/internal/config"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
)

const (
	additionalBranchFlag = "additional_branch"
	mainBranchFlag       = "main_branch"
)

// creteCmd represents the crete command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create merge requests",
	Long: `This command get required information from git and yandex tracker and ask addition information in cli for creating
merge requests from current branch to target main branch (configured) and additional branches (configured).
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

		createAdditional, _ := cmd.Flags().GetBool(additionalBranchFlag)

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

	const additionalBranchUsage = "create merge requests to additional branches"
	createCmd.Flags().BoolP(
		additionalBranchFlag, "a", false, additionalBranchUsage,
	)
}

func logMRs(mrs []gitlabcore.ResultMRInfo) {
	for _, mr := range mrs {
		log := fmt.Sprintf("\nMR to branch: %s", mr.Branch)
		if mr.URL != "" {
			log += fmt.Sprintf("\nurl: %s", mr.URL)
		}
		if mr.Err != nil {
			log += fmt.Sprintf("\nerror: %s", mr.Err.Error())
		}
		fmt.Println(log)
	}
}
