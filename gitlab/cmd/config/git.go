package configcmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	defaultMainBranch = "main"
	branchesDelim     = ","
)

func GetMainBranch(defaultVal string) (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("empty branch")
		}
		return nil
	}

	if defaultVal == "" {
		defaultVal = defaultMainBranch
	}

	prompt := promptui.Prompt{
		Label:     "Main branch",
		Default:   defaultVal,
		Validate:  validate,
		AllowEdit: true,
	}

	host, err := prompt.Run()

	return host, err
}

func GetAdditionalBranch(defaultVal []string) ([]string, error) {
	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Additional branches (release%sprepod%stest)", branchesDelim, branchesDelim),
		Default:   strings.Join(defaultVal, branchesDelim),
		AllowEdit: true,
	}

	branches, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	if branches == "" {
		return nil, nil
	}

	return strings.Split(branches, branchesDelim), err
}
