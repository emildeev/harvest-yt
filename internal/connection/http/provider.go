package httpconn

import (
	"context"
	"strconv"

	"github.com/emildeev/go-harvest/harvest"
	tracker "github.com/emildeev/yandex-tracker-go"
	"golang.org/x/oauth2"

	"github.com/emildeev/harvest-yt/internal/config"
)

type Connection struct {
	Harvest  *harvest.APIClient
	YTracker *tracker.TrackerClient
}

func New(ctx context.Context, cfg config.HTTP) (*Connection, error) {
	yTrackerClient := tracker.New("OAuth "+cfg.YTracker.Token, strconv.Itoa(cfg.YTracker.OrgID), "")
	harvestClient := harvest.NewAPIClient(
		oauth2.NewClient(
			ctx, oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: cfg.Harvest.Token,
				},
			),
		),
	)
	harvestClient.AccountID = cfg.Harvest.AccountID

	return &Connection{
		Harvest:  harvestClient,
		YTracker: yTrackerClient,
	}, nil
}
