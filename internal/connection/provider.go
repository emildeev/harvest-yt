package connection

import (
	"context"

	"github.com/emildeev/harvest-yt/internal/config"
	httpconn "github.com/emildeev/harvest-yt/internal/connection/http"
)

type Connection struct {
	HTTP *httpconn.Connection
}

func New(ctx context.Context, cfg config.Config) (*Connection, error) {
	httpConn, err := httpconn.New(ctx, cfg.HTTP)
	if err != nil {
		return nil, err
	}

	return &Connection{
		HTTP: httpConn,
	}, nil
}
