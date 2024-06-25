package mr

import (
	"context"

	"github.com/emildeev/gitlab_helper/internal/config"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"
)

type gitServiceI interface {
	GetRepoURL(_ context.Context) (url string, err error)
	GetCurrentBrunch() (brunch string, err error)
	GetTicketFromBrunch(brunch string) (string, error)
}

type gitlabServiceI interface {
	GetProjectIDByURL(ctx context.Context, url string) (projectID gitlabcore.ProjectID, err error)
	CreateMR(ctx context.Context, mr gitlabcore.MRInfo) (url string, err error)
	GetDefaultMRTemplateDescription(ctx context.Context, projectID gitlabcore.ProjectID) (description string, err error)
	FillMRTemplateDescription(_ context.Context, description, tickerURL string) string
	FindOpenedByBrunch(
		ctx context.Context, projectID gitlabcore.ProjectID, sourceBranch string,
	) (mrs []gitlabcore.CreatedMRInfo, err error)
	UpdateMRDescription(
		ctx context.Context, projectID gitlabcore.ProjectID, mrID gitlabcore.MRID, description string,
	) (url string, err error)
}

type yTrackerServiceI interface {
	GetTicket(ticketKey string) (ticket ytrackercore.Ticket, err error)
	SetMR(ticketKey string, mr string) (err error)
}

type UseCase struct {
	cfg             config.Repo
	gitlabService   gitlabServiceI
	gitService      gitServiceI
	yTrackerService yTrackerServiceI
}

func New(
	cfg config.Repo,
	gitService gitServiceI,
	gitlabService gitlabServiceI,
	yTrackerService yTrackerServiceI,
) *UseCase {
	return &UseCase{
		cfg:             cfg,
		gitService:      gitService,
		gitlabService:   gitlabService,
		yTrackerService: yTrackerService,
	}
}
