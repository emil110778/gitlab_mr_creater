package gitlab

import (
	"errors"

	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/mr"
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/project"
	projecttemplate "github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/project_template"
	"github.com/xanzy/go-gitlab"
)

type Provider struct {
	MR              *mr.Adapter
	Project         *project.Adapter
	ProjectTemplate *projecttemplate.Adapter
}

func New(gitlabClient *gitlab.Client) (*Provider, error) {
	if gitlabClient == nil {
		return nil, errors.New("gitlab is nil")
	}
	return &Provider{
		MR:              mr.New(gitlabClient.MergeRequests),
		Project:         project.New(gitlabClient.Projects),
		ProjectTemplate: projecttemplate.New(gitlabClient.ProjectTemplates),
	}, nil
}
