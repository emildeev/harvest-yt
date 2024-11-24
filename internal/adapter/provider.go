package adapter

import (
	"github.com/emildeev/harvest-yt/internal/adapter/http"
	"github.com/emildeev/harvest-yt/internal/connection"
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
