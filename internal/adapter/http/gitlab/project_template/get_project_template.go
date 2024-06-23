package projecttemplate

import (
	"context"

	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/project_template/port"
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
	httpcore "github.com/emil110778/gitlab_mr_creator/internal/core/http"
	"github.com/xanzy/go-gitlab"
)

const (
	MRTemplateMergeRequest = "merge_requests"
)

func (adapter *Adapter) GetMRTemplate(
	ctx context.Context, projectID int, templateName string,
) (mrTemplate gitlabcore.MRTemplate, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "GetProjectTemplate", mrTemplate)

	template, responseRaw, err := adapter.client.GetProjectTemplate(
		projectID, MRTemplateMergeRequest, templateName, gitlab.WithContext(ctx),
	)

	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	mrTemplate = port.GetGetMRTemplateResponse(template)

	return mrTemplate, nil
}
