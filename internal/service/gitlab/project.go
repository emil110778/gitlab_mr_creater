package gitlab

import (
	"context"
	"fmt"

	"github.com/emildeev/gitlab_helper/internal/core"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
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
	return projectID, fmt.Errorf("GetProjectIDByURL: %w", core.ErrNotFound)
}
