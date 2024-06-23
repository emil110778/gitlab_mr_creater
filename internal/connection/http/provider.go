package httpconn

import (
	"strconv"

	tracker "github.com/dvsnin/yandex-tracker-go"
	"github.com/emil110778/gitlab_mr_creator/internal/config"
	"github.com/xanzy/go-gitlab"
)

type Connection struct {
	Gitlab   *gitlab.Client
	YTracker tracker.Client
}

func New(cfg config.HTTP) (*Connection, error) {
	gitlabClient, err := gitlab.NewClient(cfg.Gitlab.Token, gitlab.WithBaseURL(cfg.Gitlab.Host))
	if err != nil {
		return nil, err
	}

	yTrackerClient := tracker.New("OAuth "+cfg.YTracker.Token, strconv.Itoa(cfg.YTracker.OrgID), "")

	return &Connection{
		Gitlab:   gitlabClient,
		YTracker: yTrackerClient,
	}, nil
}
