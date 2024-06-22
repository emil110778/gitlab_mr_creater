package httpconn

import (
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/xanzy/go-gitlab"
)

type Connection struct {
	Gitlab *gitlab.Client
}

func New(connectionProvider config.HTTP) (*Connection, error) {
	client, err := gitlab.NewClient(connectionProvider.Gitlab.Token, gitlab.WithBaseURL(connectionProvider.Gitlab.Host))
	if err != nil {
		return nil, err
	}
	return &Connection{
		Gitlab: client,
	}, nil
}
