package mr

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	httpcore "github.com/emildeev/gitlab_helper/internal/core/http"
	ytrackercore "github.com/emildeev/gitlab_helper/internal/core/y_tracker"
	"github.com/emildeev/gitlab_helper/pkg/helper"
)

func (uc *UseCase) Create(ctx context.Context, createAdditional bool) (mrs []gitlabcore.ResultMRInfo, err error) {
	errorWrapper := func(err error) ([]gitlabcore.ResultMRInfo, error) {
		return mrs, fmt.Errorf("MR UseCase: Create: %w", err)
	}
	repoURL, err := uc.gitService.GetRepoURL(ctx)
	if err != nil {
		return errorWrapper(err)
	}

	slog.Debug("mr uc Create:", "url", repoURL)

	projectID, err := uc.gitlabService.GetProjectIDByURL(ctx, repoURL)
	if err != nil {
		return errorWrapper(err)
	}

	slog.Debug("mr uc Create:", "projectID", projectID)

	currentBranch, err := uc.gitService.GetCurrentBranch()
	if err != nil {
		return errorWrapper(err)
	}

	slog.Debug("mr uc Create:", "currentBranch", currentBranch)

	ticket, err := uc.getMRTicket(ctx, currentBranch)
	if err != nil {
		slog.Warn("getMRTicket:", "err", err)
	}

	slog.Debug("mr uc Create:", "ticket.Title", ticket.Title, "ticket.Key", ticket.Key)

	var title string
	if ticket.Key != "" && ticket.Title != "" {
		title = fmt.Sprintf("%s: %s", ticket.Key, ticket.Title)
	}

	mrs = uc.createMRs(ctx, projectID, currentBranch, title, createAdditional, ticket.Key)

	slog.Debug("mr uc Create:", "mrs", mrs)

	if len(mrs) != 0 && title != "" {
		for _, mr := range mrs {
			if mr.Branch == uc.cfg.MainBranch && mr.URL != "" {
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
	ctx context.Context, projectID gitlabcore.ProjectID, currentBranch, title string,
	createAdditional bool, ticketKey string,
) (mrs []gitlabcore.ResultMRInfo) {
	createdMrs, err := uc.gitlabService.FindOpenedByBranch(ctx, projectID, currentBranch)
	if err != nil {
		slog.Warn("FindOpenedByBranch:", "err", err)
	}

	createdMrsMap := helper.GetMapFromSliceByField(createdMrs, func(obj gitlabcore.CreatedMRInfo) string {
		return obj.TargetBranch
	})

	var mainDescription string
	if _, exist := createdMrsMap[uc.cfg.MainBranch]; !exist {
		mainDescription, err = uc.gitlabService.GetDefaultMRTemplateDescription(ctx, projectID)
		if err != nil {
			slog.Warn("GetDefaultMRTemplateDescription:", "err", err)
		}

		if mainDescription != "" {
			mainDescription = uc.gitlabService.FillMRTemplateDescription(ctx, mainDescription, ticketKey)
		}
	}

	mainMr := uc.createMR(ctx, projectID, currentBranch, uc.cfg.MainBranch, title, gitlabcore.MROptionalInfo{
		Draft:                true,
		ApprovalsBeforeMerge: helper.GetPointer(2),
		Description:          helper.GetPointer(mainDescription),
		RemoveSourceBranch:   helper.GetPointer(true),
	}, createdMrsMap)

	mrs = append(mrs, mainMr)

	if createAdditional {
		for _, additionalBranch := range uc.cfg.AdditionalBranches {
			mrs = append(mrs, uc.createMR(
				ctx, projectID, currentBranch, additionalBranch, title, gitlabcore.MROptionalInfo{
					Description: helper.GetPointer(mainMr.URL),
				}, createdMrsMap,
			))
		}
	}
	return mrs
}

func (uc *UseCase) getMRTicket(ctx context.Context, branch string) (ticket ytrackercore.Ticket, err error) {
	ticketKey, err := uc.gitService.GetTicketFromBranch(branch)
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
	ctx context.Context, projectID gitlabcore.ProjectID, currentBranch, targetBranch string, title string,
	optionalInfo gitlabcore.MROptionalInfo,
	createdMrsMap map[string]gitlabcore.CreatedMRInfo,
) (mr gitlabcore.ResultMRInfo) {
	if mrCreated, exist := createdMrsMap[targetBranch]; exist {
		mr.URL = mrCreated.URL
		mr.Branch = mrCreated.TargetBranch
		mr.Err = errors.New("mr already exists")

		if optionalInfo.Description != nil && *optionalInfo.Description != "" {
			_, err := uc.gitlabService.UpdateMRDescription(ctx, projectID, mrCreated.ID, *optionalInfo.Description)
			if err != nil {
				slog.Error("UpdateMRDescription:", "err", err)
			}
		}
	} else {
		mrURL, err := uc.gitlabService.CreateMR(ctx, gitlabcore.MRInfo{
			Title:          title,
			ProjectID:      projectID,
			SourceBranch:   currentBranch,
			TargetBranch:   targetBranch,
			MROptionalInfo: optionalInfo,
		})
		if err != nil {
			slog.Error("CreateMR:", "err", err)
		}
		mr = gitlabcore.ResultMRInfo{
			Branch: targetBranch,
			URL:    mrURL,
			Err:    getMRError(err),
		}
	}
	return mr
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
