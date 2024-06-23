package connection

import (
	"github.com/emildeev/gitlab_helper/internal/config"
	httpconn "github.com/emildeev/gitlab_helper/internal/connection/http"
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
