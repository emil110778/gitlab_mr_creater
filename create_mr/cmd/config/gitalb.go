package configcmd

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/manifoldco/promptui"
)

const (
	defaultGitlabHost = "https://gitlab.com"
)

func GitlabGetHost(defaultVal string) (string, error) {
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
		Label:     "Gitlab host",
		Default:   defaultVal,
		Validate:  validate,
		AllowEdit: true,
	}

	host, err := prompt.Run()

	return host, err
}

func GitlabGetToken(defaultVal string) (string, error) {
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
		Label:    "Gitlab token",
		Default:  defaultVal,
		Mask:     '*',
		Validate: validate,
	}

	token, err := prompt.Run()

	return token, err
}
