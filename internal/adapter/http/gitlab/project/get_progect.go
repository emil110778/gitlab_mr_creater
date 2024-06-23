package project

import (
	"context"

	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/project/port"
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
	httpcore "github.com/emil110778/gitlab_mr_creator/internal/core/http"
	"github.com/xanzy/go-gitlab"
)

func (adapter *Adapter) GetProjects(ctx context.Context) (gitlabProjects []gitlabcore.Project, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "GetProjects", gitlabProjects)

	req := port.GetGetProjectsRequest()

	projects, responseRaw, err := adapter.client.ListProjects(req, gitlab.WithContext(ctx))

	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	gitlabProjects = port.GetGetProjectsResponse(projects)

	return gitlabProjects, nil
}
