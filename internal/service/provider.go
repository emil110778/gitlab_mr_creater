package service

import (
	"github.com/emil110778/gitlab_mr_creator/internal/adapter"
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/emil110778/gitlab_mr_creator/internal/service/git"
	"github.com/emil110778/gitlab_mr_creator/internal/service/gitlab"
)

type Provider struct {
	Git    *git.Service
	Gitlab *gitlab.Service
}

func New(_ config.Config, provider adapter.Provider) *Provider {
	return &Provider{
		Git:    git.New(),
		Gitlab: gitlab.New(provider.HTTP.Gitlab.MR, provider.HTTP.Gitlab.Project),
	}
}
