package adapter

import (
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http"
	"github.com/emil110778/gitlab_mr_creator/internal/connection"
)

type Provider struct {
	HTTP *http.Provider
}

func New(connections *connection.Connection) (Provider, error) {
	httpProvider, err := http.New(connections.HTTP)
	if err != nil {
		return Provider{}, err
	}

	return Provider{
		HTTP: httpProvider,
	}, nil
}
