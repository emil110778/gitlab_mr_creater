package mr

import (
	"context"

	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/mr/port"
	"github.com/emil110778/gitlab_mr_creator/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
	"github.com/emil110778/gitlab_mr_creator/internal/core/http"
	"github.com/xanzy/go-gitlab"
)

func (adapter *Adapter) CreateMR(
	ctx context.Context, mr gitlabcore.MRInfo,
) (url string, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "CreateMR", url)

	gitlabMR := port.GetMRRequest(mr)

	resultMR, responseRaw, err := adapter.client.CreateMergeRequest(int(mr.ProjectID), &gitlabMR, gitlab.WithContext(ctx))
	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	url = port.GetMRResponse(resultMR)

	return url, nil
}
