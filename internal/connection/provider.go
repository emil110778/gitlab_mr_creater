package connection

import (
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	httpconn "github.com/emil110778/gitlab_mr_creator/internal/connection/http"
)

type Connection struct {
	HTTP *httpconn.Connection
}

func New(cfg config.Config) (*Connection, error) {
	httpConn, err := httpconn.New(cfg.HTTP)
	if err != nil {
		return nil, err
	}

	return &Connection{
		HTTP: httpConn,
	}, nil
}
