package mr

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/emildeev/gitlab_helper/internal/config"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	httpcore "github.com/emildeev/gitlab_helper/internal/core/http"
	ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"
	"github.com/emildeev/gitlab_helper/pkg/helper"
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

func (uc *UseCase) Create(ctx context.Context) (mrs []gitlabcore.CreatedMRInfo, err error) {
	errorWrapper := func(err error) ([]gitlabcore.CreatedMRInfo, error) {
		return mrs, fmt.Errorf("MR UseCase: Create: %w", err)
	}
	repoURL, err := uc.gitService.GetRepoURL(ctx)
	if err != nil {
		return errorWrapper(err)
	}

	projectID, err := uc.gitlabService.GetProjectIDByURL(ctx, repoURL)
	if err != nil {
		return errorWrapper(err)
	}

	currentBrunch, err := uc.gitService.GetCurrentBrunch()
	if err != nil {
		return errorWrapper(err)
	}

	ticket, err := uc.getMRTitle(ctx, currentBrunch)
	if err != nil {
		slog.Warn("getMRTitle error: ", err)
	}
	var title string
	if ticket.Key != "" && ticket.Title != "" {
		title = fmt.Sprintf("%s: %s", ticket.Key, ticket.Title)
	}

	getMRTemplateDescription, err := uc.gitlabService.GetDefaultMRTemplateDescription(ctx, projectID)
	if err != nil {
		slog.Warn("GetDefaultMRTemplateDescription error: ", err)
	}

	if getMRTemplateDescription != "" {
		getMRTemplateDescription = uc.gitlabService.FillMRTemplateDescription(ctx, getMRTemplateDescription, ticket.Key)
	}

	mrs, err = uc.createMRs(ctx, projectID, currentBrunch, title, getMRTemplateDescription)
	if err != nil {
		return errorWrapper(err)
	}

	if len(mrs) != 0 && title != "" {
		for _, mr := range mrs {
			if mr.Brunch == uc.cfg.MainBrunch && mr.URL != "" {
				err = uc.yTrackerService.SetMR(ticket.Key, mr.URL)
				if err != nil {
					return errorWrapper(err)
				}
			}
		}
	}

	return
}

func (uc *UseCase) createMRs(
	ctx context.Context, projectID gitlabcore.ProjectID, currentBrunch, title, mainDescription string,
) (mrs []gitlabcore.CreatedMRInfo, err error) {
	mainMr := uc.createMR(ctx, projectID, currentBrunch, uc.cfg.MainBrunch, title, gitlabcore.MROptionalInfo{
		Draft:                true,
		ApprovalsBeforeMerge: helper.GetPointer(2),
		Description:          helper.GetPointer(mainDescription),
	})
	mrs = append(mrs, mainMr)

	for _, additionalBrunch := range uc.cfg.AdditionalBrunches {
		mrs = append(mrs, uc.createMR(
			ctx, projectID, currentBrunch, additionalBrunch, title, gitlabcore.MROptionalInfo{
				Description: helper.GetPointer(mainMr.URL),
			},
		))
	}
	return mrs, nil
}

func (uc *UseCase) getMRTitle(ctx context.Context, brunch string) (ticket ytrackercore.Ticket, err error) {
	ticketKey, err := uc.gitService.GetTicketFromBrunch(brunch)
	if err != nil {
		return ticket, fmt.Errorf("getTaskTitle: %w", err)
	}
	ticket, err = uc.yTrackerService.GetTicket(ticketKey)
	if err != nil {
		return ticket, fmt.Errorf("getTaskTitle: %w", err)
	}

	return ticket, nil
}

func (uc *UseCase) createMR(
	ctx context.Context, projectID gitlabcore.ProjectID, currentBrunch, targetBranch string, title string,
	optionalInfo gitlabcore.MROptionalInfo,
) (mr gitlabcore.CreatedMRInfo) {
	mrUrl, err := uc.gitlabService.CreateMR(ctx, gitlabcore.MRInfo{
		Title:          title,
		ProjectID:      projectID,
		SourceBranch:   currentBrunch,
		TargetBranch:   targetBranch,
		MROptionalInfo: optionalInfo,
	})
	if err != nil {
		slog.Debug("CreateMR error: ", err)
	}
	return gitlabcore.CreatedMRInfo{
		Brunch: targetBranch,
		URL:    mrUrl,
		Err:    getMRError(err),
	}
}

func getMRError(err error) error {
	if err == nil {
		return nil
	}
	var errHTTP *httpcore.HTTPError
	ok := errors.As(err, &errHTTP)
	if ok {
		if errHTTP.StatusCode == http.StatusConflict {
			return errors.New("MR already exists")
		}
		if errHTTP.StatusCode == http.StatusForbidden {
			return errors.New("you cant access for create mr")
		}
	}
	return errors.New("mr not created")
}
