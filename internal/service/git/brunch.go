package git

import (
	"fmt"
	"strings"
)

func (service *Service) GetCurrentBrunch() (brunch string, err error) {
	errHandler := func(err error) (string, error) {
		return brunch, fmt.Errorf("GetCurrentBrunch: %w", err)
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

func (service *Service) GetTicketFromBrunch(brunch string) (string, error) {
	keyWithDelim := service.regExpTaskKeyWithDelim.FindString(brunch)
	key := service.regExpTaskKey.FindString(keyWithDelim)
	if key == "" {
		return "", fmt.Errorf("ticket not found in brunch name %s", brunch)
	}
	return key, nil
}
