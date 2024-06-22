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
