package mr

import (
	"context"

	"github.com/xanzy/go-gitlab"

	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/mr/port"
	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	httpcore "github.com/emildeev/gitlab_helper/internal/core/http"
)

func (adapter *Adapter) List(
	ctx context.Context, filter gitlabcore.FilterMR,
) (mrs []gitlabcore.CreatedMRInfo, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "FindByBranches", mrs)

	options := port.GetListMRRequest(filter)

	resultMR, responseRaw, err := adapter.client.ListProjectMergeRequests(
		int(filter.ProjectID), options, gitlab.WithContext(ctx),
	)
	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	mrs = port.GetListMRResponse(resultMR)

	return mrs, nil
}
