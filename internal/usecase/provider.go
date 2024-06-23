package usecase

import (
	"github.com/emildeev/gitlab_helper/internal/config"
	"github.com/emildeev/gitlab_helper/internal/service"
	"github.com/emildeev/gitlab_helper/internal/usecase/mr"
)

type Provider struct {
	MR *mr.UseCase
}

func New(cfg config.Config, provider *service.Provider) (*Provider, error) {
	return &Provider{
		MR: mr.New(cfg.Repo, provider.Git, provider.Gitlab, provider.YTracker),
	}, nil
}
