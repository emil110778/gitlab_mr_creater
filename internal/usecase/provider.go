package usecase

import (
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/emil110778/gitlab_mr_creator/internal/service"
	"github.com/emil110778/gitlab_mr_creator/internal/usecase/mr"
)

type Provider struct {
	MR *mr.UseCase
}

func New(cfg config.Config, provider *service.Provider) (*Provider, error) {
	return &Provider{
		MR: mr.New(cfg.Repo, provider.Git, provider.Gitlab),
	}, nil
}
