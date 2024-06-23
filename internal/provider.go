package internal

import (
	"github.com/emildeev/gitlab_helper/internal/adapter"
	"github.com/emildeev/gitlab_helper/internal/config"
	"github.com/emildeev/gitlab_helper/internal/connection"
	"github.com/emildeev/gitlab_helper/internal/service"
	"github.com/emildeev/gitlab_helper/internal/usecase"
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
	serviceProvider, err := service.New(cfg, drivenProvider)
	if err != nil {
		return provider, err
	}
	useCaseProvider, err := usecase.New(cfg, serviceProvider)
	if err != nil {
		return provider, err
	}
	return useCaseProvider, nil
}
