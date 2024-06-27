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

	tracker "github.com/dvsnin/yandex-tracker-go"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"

	configcmd "github.com/emildeev/gitlab_helper/gitlab/cmd/config"
	"github.com/emildeev/gitlab_helper/internal/config"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure environment for tool",
	Long: `This command will configure environment for tool:
gitlab credentials
yandex tracker credentials
and repository brunch configuration
`,
	SilenceErrors: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		currentCfg := config.NewWithoutValidate()

		host, err := configcmd.GitlabGetHost(currentCfg.HTTP.Gitlab.Host)
		if err != nil {
			return err
		}

		token, err := configcmd.GitlabGetToken(currentCfg.HTTP.Gitlab.Token)
		if err != nil {
			return err
		}

		client, err := gitlab.NewClient(token, gitlab.WithBaseURL(host))
		if err != nil {
			return err
		}
		_, resp, err := client.Version.GetVersion()
		if err != nil {
			if resp != nil && resp.StatusCode == http.StatusUnauthorized {
				slog.Error("gitlab authorization error", err)
				return errors.New("gitlab authorization error")
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
			if resp.StatusCode == http.StatusUnauthorized {
				slog.Error("yandex tracker authorization error", err)
				return errors.New("yandex tracker authorization error")
			}
			return err
		}

		mainBrunch, err := configcmd.GetMainBrunch(currentCfg.Repo.MainBrunch)
		if err != nil {
			return err
		}

		additionalBrunches, err := configcmd.GetAdditionalBrunch(currentCfg.Repo.AdditionalBrunches)
		if err != nil {
			return err
		}

		cfg := config.Config{
			HTTP: config.HTTP{
				Gitlab: config.Gitlab{
					Host:  host,
					Token: token,
				},
				YTracker: config.YTracker{
					Token: yTrackerToken,
					OrgID: yTrackerOrgID,
				},
			},
			Repo: config.Repo{
				MainBrunch:         mainBrunch,
				AdditionalBrunches: additionalBrunches,
			},
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
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
