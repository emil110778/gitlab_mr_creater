package gitlab

import (
	"context"
	"fmt"

	"github.com/emil110778/gitlab_mr_creator/internal/core"
	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
)

func (service *Service) GetProjectIDByURL(ctx context.Context, url string) (projectID gitlabcore.ProjectID, err error) {
	gitlabProjects, err := service.projectAdapter.GetProjects(ctx)
	if err != nil {
		return projectID, fmt.Errorf("GetProjectIDByURL: %w", err)
	}
	for _, project := range gitlabProjects {
		if project.URL == url {
			return project.ID, nil
		}
	}
	return projectID, core.ErrNotFound
}
