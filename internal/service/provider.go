package service

import (
	"github.com/emil110778/gitlab_mr_creator/internal/adapter"
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/emil110778/gitlab_mr_creator/internal/service/git"
	"github.com/emil110778/gitlab_mr_creator/internal/service/gitlab"
	ytracker "github.com/emil110778/gitlab_mr_creator/internal/service/y_tracker"
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
