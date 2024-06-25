package gitlab

import (
	"context"
	"fmt"

	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/emildeev/gitlab_helper/pkg/helper"
)

func (service *Service) FindOpenedByBrunch(
	ctx context.Context, projectID gitlabcore.ProjectID, sourceBranch string,
) (mrs []gitlabcore.CreatedMRInfo, err error) {
	mrs, err = service.GetListMR(ctx, gitlabcore.FilterMR{
		ProjectID:    projectID,
		SourceBranch: helper.GetPointer(sourceBranch),
		State:        helper.GetPointer(gitlabcore.MRStateOpened),
	})
	if err != nil {
		return mrs, fmt.Errorf("FindOpenedByBrunch: %w", err)
	}
	return mrs, nil
}

func (service *Service) GetListMR(
	ctx context.Context, filter gitlabcore.FilterMR) (mrs []gitlabcore.CreatedMRInfo, err error) {
	mrs, err = service.mrAdapter.List(ctx, filter)
	if err != nil {
		return mrs, fmt.Errorf("GetListMR: %w", err)
	}
	return mrs, nil
}

func (service *Service) CreateMR(ctx context.Context, mr gitlabcore.MRInfo) (url string, err error) {
	mrID, err := service.mrAdapter.Create(ctx, mr)
	if err != nil {
		return url, fmt.Errorf("CreateMR: %w", err)
	}
	return mrID, nil
}

func (service *Service) UpdateMRDescription(
	ctx context.Context, projectID gitlabcore.ProjectID, mrID gitlabcore.MRID, description string,
) (url string, err error) {
	url, err = service.UpdateMR(ctx, gitlabcore.MRUpdateInfo{
		ID:        mrID,
		ProjectID: projectID,
		MROptionalUpdateInfo: gitlabcore.MROptionalUpdateInfo{
			Description: helper.GetPointer(description),
		},
	})
	if err != nil {
		return url, fmt.Errorf("UpdateMR: %w", err)
	}
	return url, nil
}

func (service *Service) UpdateMR(ctx context.Context, mr gitlabcore.MRUpdateInfo) (url string, err error) {
	url, err = service.mrAdapter.Update(ctx, mr)
	if err != nil {
		return url, fmt.Errorf("UpdateMR: %w", err)
	}
	return url, nil
}
