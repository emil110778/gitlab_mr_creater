package git

import (
	"fmt"
	"strings"
)

func (service *Service) CheckBrunch(branch string) bool {
	repo, err := getRepo()
	if err != nil {
		return false
	}

	branchCfg, err := repo.Branch(branch)
	if err != nil {
		return false
	}
	err = branchCfg.Validate()
	if err != nil {
		return false
	}

	return true
}

func (service *Service) GetCurrentBranch() (branch string, err error) {
	errHandler := func(err error) (string, error) {
		return branch, fmt.Errorf("GetCurrentBranch: %w", err)
	}

	repo, err := getRepo()
	if err != nil {
		return errHandler(err)
	}

	h, err := repo.Head()
	if err != nil {
		return errHandler(err)
	}
	return strings.TrimPrefix(string(h.Name()), "refs/heads/"), nil
}

func (service *Service) GetTicketFromBranch(branch string) (string, error) {
	keyWithDelim := service.regExpTaskKeyWithDelim.FindString(branch)
	key := service.regExpTaskKey.FindString(keyWithDelim)
	if key == "" {
		return "", fmt.Errorf("ticket not found in branch name %s", branch)
	}
	return key, nil
}
