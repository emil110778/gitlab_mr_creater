package port

import (
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/xanzy/go-gitlab"
)

func GetGetProjectsRequest() *gitlab.ListProjectsOptions {
	return &gitlab.ListProjectsOptions{}
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
