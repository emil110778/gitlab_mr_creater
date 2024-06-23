package http

import (
	"errors"
	"fmt"

	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab"
	ytracker "github.com/emil110778/gitlab_mr_creator/internal/adapter/http/y_tracker"
	httpconn "github.com/emil110778/gitlab_mr_creator/internal/connection/http"
)

type Provider struct {
	Gitlab   *gitlab.Provider
	YTracker *ytracker.Adapter
}

func New(connections *httpconn.Connection) (*Provider, error) {
	errorWrapper := func(err error) (*Provider, error) {
		return nil, fmt.Errorf("error configure http adapter %w", err)
	}

	if connections == nil {
		return errorWrapper(errors.New("http connections is nil"))
	}

	gitlabProvider, err := gitlab.New(connections.Gitlab)
	if err != nil {
		return errorWrapper(err)
	}

	yTrackerAdapter := ytracker.New(connections.YTracker)

	return &Provider{
		Gitlab:   gitlabProvider,
		YTracker: yTrackerAdapter,
	}, nil
}
