package gitlab

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/emildeev/gitlab_helper/internal/core"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
)

func (service *Service) GetProjectIDByURL(
	ctx context.Context,
	urlStr string,
) (projectID gitlabcore.ProjectID, err error) {
	var projectName *string
	urlParsed, err := url.Parse(urlStr)
	if err == nil && urlParsed != nil {
		splitPath := strings.Split(urlParsed.Path, "/")
		projectName = &splitPath[len(splitPath)-1]
	}

	gitlabProjects, err := service.projectAdapter.GetProjectsByProjectName(ctx, projectName)
	if err != nil {
		return projectID, fmt.Errorf("GetProjectIDByURL: %w", err)
	}
	for _, project := range gitlabProjects {
		if project.URL == urlStr {
			return project.ID, nil
		}
	}
	return projectID, fmt.Errorf("GetProjectIDByURL: %w", core.ErrNotFound)
}
