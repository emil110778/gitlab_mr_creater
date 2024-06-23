package mr

import (
	"context"

	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/mr/port"
	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/emildeev/gitlab_helper/internal/core/http"
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
