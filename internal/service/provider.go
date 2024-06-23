package service

import (
	"github.com/emildeev/gitlab_helper/internal/adapter"
	"github.com/emildeev/gitlab_helper/internal/config"
	"github.com/emildeev/gitlab_helper/internal/service/git"
	"github.com/emildeev/gitlab_helper/internal/service/gitlab"
	ytracker "github.com/emildeev/gitlab_helper/internal/service/y_tracker"
)

type Provider struct {
	Git      *git.Service
	Gitlab   *gitlab.Service
	YTracker *ytracker.Service
}

func New(_ config.Config, provider adapter.Provider) (*Provider, error) {
	gitService, err := git.New()
	if err != nil {
		return nil, err
	}
	gitLabService, err := gitlab.New(
		provider.HTTP.Gitlab.MR, provider.HTTP.Gitlab.Project, provider.HTTP.Gitlab.ProjectTemplate,
	)
	if err != nil {
		return nil, err
	}
	return &Provider{
		Git:      gitService,
		Gitlab:   gitLabService,
		YTracker: ytracker.New(provider.HTTP.YTracker),
	}, nil
}
