package project

import (
	"context"

	"github.com/xanzy/go-gitlab"

	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/project/port"
	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	httpcore "github.com/emildeev/gitlab_helper/internal/core/http"
)

func (adapter *Adapter) GetProjects(ctx context.Context) (gitlabProjects []gitlabcore.Project, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "GetProjects", gitlabProjects)

	req := port.GetGetProjectsRequest(nil)

	projects, responseRaw, err := adapter.client.ListProjects(req, gitlab.WithContext(ctx))

	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	gitlabProjects = port.GetGetProjectsResponse(projects)

	return gitlabProjects, nil
}

func (adapter *Adapter) GetProjectsByProjectName(
	ctx context.Context,
	projectName *string,
) (gitlabProjects []gitlabcore.Project, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "GetProjects", gitlabProjects)

	req := port.GetGetProjectsRequest(projectName)

	projects, responseRaw, err := adapter.client.ListProjects(req, gitlab.WithContext(ctx))

	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	gitlabProjects = port.GetGetProjectsResponse(projects)

	return gitlabProjects, nil
}
