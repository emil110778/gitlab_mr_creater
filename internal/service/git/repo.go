package git

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
)

func (service *Service) GetRepoURL(_ context.Context) (url string, err error) {
	errHandler := func(err error) (string, error) {
		return url, fmt.Errorf("GetCurrentBrunch: %w", err)
	}

	repo, err := getRepo()
	if err != nil {
		return errHandler(err)
	}

	remotes, err := repo.Remotes()
	if err != nil {
		return errHandler(err)
	}

	if len(remotes) == 0 {
		return errHandler(fmt.Errorf("no remotes found"))
	}

	urls := remotes[0].Config().URLs

	if len(urls) == 0 {
		return errHandler(fmt.Errorf("no remotes urls found"))
	}

	url = strings.TrimSuffix(urls[0], ".git")

	return url, nil
}

func getRepo() (*git.Repository, error) {
	dir, err := getExecDir()
	if err != nil {
		return nil, err
	}

	if dir == "" {
		return nil, fmt.Errorf("error getting current directory")
	}

	repo, err := findRepo(dir)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func findRepo(dir string) (*git.Repository, error) {
	var dirs []string
	var repo *git.Repository
	var err error

	for dir != "" && dir != "/" {
		repo, err = git.PlainOpen(dir)
		if err != nil {
			dirs = append(dirs, dir)
			dir = path.Dir(dir)
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("error getting git repository in paths: %v, %w", dirs, err)
	}
	return repo, nil
}

func getExecDir() (string, error) {
	return os.Getwd()
}
