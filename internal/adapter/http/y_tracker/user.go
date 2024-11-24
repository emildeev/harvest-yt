package ytracker

import (
	"context"

	"github.com/emildeev/harvest-yt/internal/adapter/http/y_tracker/port"
	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func (adapter *Adapter) GetMyself(_ context.Context) (myself ytrackercore.User, err error) {
	user, err := adapter.client.Myself()
	if err != nil {
		return myself, err
	}

	return port.UserToCore(user), nil
}
