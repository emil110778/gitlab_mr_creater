package configcmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	defaultMainBrunch = "main"
	branchesDelim     = ","
)

func GetMainBrunch(defaultVal string) (string, error) {
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
		Label:     "Main brunch",
		Default:   defaultVal,
		Validate:  validate,
		AllowEdit: true,
	}

	host, err := prompt.Run()

	return host, err
}

func GetAdditionalBrunch(defaultVal []string) ([]string, error) {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Additional brunches (release%sprepod%stest)", branchesDelim, branchesDelim),
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
