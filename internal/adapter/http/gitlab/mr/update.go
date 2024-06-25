package mr

import (
	"context"

	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/mr/port"
	"github.com/emildeev/gitlab_helper/internal/adapter/http/gitlab/response"
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/emildeev/gitlab_helper/internal/core/http"
	"github.com/xanzy/go-gitlab"
)

func (adapter *Adapter) Update(
	ctx context.Context, mr gitlabcore.MRUpdateInfo,
) (url string, err error) {
	errHandleFunc := httpcore.GetHandleErrorFunc("gitlab", "CreateMR", url)

	gitlabMR := port.UpdateMRRequest(mr)

	resultMR, responseRaw, err := adapter.client.UpdateMergeRequest(
		int(mr.ProjectID), int(mr.ID), &gitlabMR, gitlab.WithContext(ctx),
	)
	resp := response.GetResponse(responseRaw)
	if httpcore.HandleHTTPError(err, resp) != nil {
		return errHandleFunc(err, resp)
	}

	url = port.GetMRResponseURL(resultMR)

	return url, nil
}
