/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

const (
	defaultGitlabHost = "https://gitlab.com"
	defaultMainBrunch = "main"
	branchesDelim     = ","
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure environment for tool",
	Long: `This command will configure environment for tool:
gitlab credentials
yandex tracker credentials
`,
	SilenceErrors: false,
	RunE: func(cmd *cobra.Command, args []string) error {
		currentCfg := config.NewWithoutValidate()

		host, err := getHost(currentCfg.HTTP.Gitlab.Host)
		if err != nil {
			return err
		}

		token, err := getToken(currentCfg.HTTP.Gitlab.Token)
		if err != nil {
			return err
		}

		client, err := gitlab.NewClient(token, gitlab.WithBaseURL(host))
		if err != nil {
			return err
		}
		version, resp, err := client.Version.GetVersion()
		if err != nil {
			if resp.StatusCode == http.StatusUnauthorized {
				return fmt.Errorf("authorization error: %w", err)
			}
			return err
		}

		fmt.Println("your Gitlab version: ", version)

		mainBrunch, err := getMainBrunch(currentCfg.Repo.MainBrunch)
		if err != nil {
			return err
		}

		additionalBrunches, err := getAdditionalBrunch(currentCfg.Repo.AdditionalBrunches)
		if err != nil {
			return err
		}

		cfg := config.Config{
			HTTP: config.HTTP{
				Gitlab: config.Gitlab{
					Host:  host,
					Token: token,
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

func getHost(defaultVal string) (string, error) {
	validate := func(input string) error {
		_, err := url.Parse(input)
		if err != nil {
			return fmt.Errorf("invalid url: %w", err)
		}
		return nil
	}

	if defaultVal == "" {
		defaultVal = defaultGitlabHost
	}

	prompt := promptui.Prompt{
		Label:     "Gitlab host:",
		Default:   defaultVal,
		Validate:  validate,
		AllowEdit: true,
	}

	host, err := prompt.Run()

	return host, err
}

func getToken(defaultVal string) (string, error) {
	validate := func(input string) error {
		validateRegexp := "[0-9a-zA-Z\\-]{20}"
		matched, err := regexp.MatchString(validateRegexp, input)
		if err != nil {
			return fmt.Errorf("validate error: %w", err)
		}
		if !matched {
			return fmt.Errorf("invalid token format, shuld be: %s", validateRegexp)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Gitlab token:",
		Default:   defaultVal,
		Mask:      '*',
		Validate:  validate,
		AllowEdit: true,
	}

	token, err := prompt.Run()

	return token, err
}

func getMainBrunch(defaultVal string) (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("empty brunch")
		}
		return nil
	}

	if defaultVal == "" {
		defaultVal = defaultMainBrunch
	}

	prompt := promptui.Prompt{
		Label:     "Main brunch:",
		Default:   defaultVal,
		Validate:  validate,
		AllowEdit: true,
	}

	host, err := prompt.Run()

	return host, err
}

func getAdditionalBrunch(defaultVal []string) ([]string, error) {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Additional brunches (a%sb%sc):", branchesDelim, branchesDelim),
		Default:   strings.Join(defaultVal, branchesDelim),
		AllowEdit: true,
	}

	brunches, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	if brunches == "" {
		return nil, nil
	}

	return strings.Split(brunches, branchesDelim), err
}
