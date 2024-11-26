package port

import (
	"math"

	"github.com/xanzy/go-gitlab"

	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
)

func GetGetProjectsRequest(projectName *string) *gitlab.ListProjectsOptions {
	return &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: math.MaxInt,
		},
		Search: projectName,
	}
}

func GetGetProjectsResponse(projects []*gitlab.Project) (gitlabProjects []gitlabcore.Project) {
	gitlabProjects = make([]gitlabcore.Project, len(projects))
	for i, project := range projects {
		gitlabProjects[i] = gitlabcore.Project{
			ID:  gitlabcore.ProjectID(project.ID),
			URL: project.WebURL,
		}
	}
	return gitlabProjects
}
