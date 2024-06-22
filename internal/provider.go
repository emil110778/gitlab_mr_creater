package internal

import (
	"github.com/emil110778/gitlab_mr_creator/internal/adapter"
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/emil110778/gitlab_mr_creator/internal/connection"
	"github.com/emil110778/gitlab_mr_creator/internal/service"
	"github.com/emil110778/gitlab_mr_creator/internal/usecase"
)

func New(cfg config.Config) (provider *usecase.Provider, err error) {
	connectionProvider, err := connection.New(cfg)
	if err != nil {
		return provider, err
	}
	drivenProvider, err := adapter.New(connectionProvider)
	if err != nil {
		return provider, err
	}
	serviceProvider := service.New(cfg, drivenProvider)
	if err != nil {
		return provider, err
	}
	useCaseProvider, err := usecase.New(cfg, serviceProvider)
	if err != nil {
		return provider, err
	}
	return useCaseProvider, nil
}
