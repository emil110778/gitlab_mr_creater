package port

import (
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/emildeev/gitlab_helper/pkg/helper"
	"github.com/xanzy/go-gitlab"
)

func GetListMRRequest(filter gitlabcore.FilterMR) *gitlab.ListProjectMergeRequestsOptions {
	return &gitlab.ListProjectMergeRequestsOptions{
		SourceBranch: helper.CopyPointer(filter.SourceBranch),
		TargetBranch: helper.CopyPointer(filter.TargetBranch),
		State:        (*string)(helper.CopyPointer(filter.State)),
	}
}

func GetListMRResponse(response []*gitlab.MergeRequest) (mrs []gitlabcore.CreatedMRInfo) {
	mrs = make([]gitlabcore.CreatedMRInfo, 0, len(response))
	for _, mr := range response {
		if mr == nil {
			continue
		}
		mrs = append(mrs, GetMRResponse(*mr))
	}
	return mrs
}
