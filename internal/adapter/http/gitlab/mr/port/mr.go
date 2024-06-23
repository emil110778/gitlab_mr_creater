package port

import (
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
	"github.com/emildeev/gitlab_helper/pkg/helper"
	"github.com/xanzy/go-gitlab"
)

const (
	draftPrefix = "Draft: "
)

func GetMRRequest(mr gitlabcore.MRInfo) gitlab.CreateMergeRequestOptions {
	var (
		assigneeIDs *[]int
		reviewerIDs *[]int
	)

	if mr.Draft {
		mr.Title = draftPrefix + mr.Title
	}

	if len(mr.Assignees) != 0 {
		assigneeIDs = helper.GetPointer(helper.SliceToInt(mr.Assignees))
	}

	if len(mr.Reviewers) != 0 {
		reviewerIDs = helper.GetPointer(helper.SliceToInt(mr.Reviewers))
	}

	return gitlab.CreateMergeRequestOptions{
		Title:                helper.GetPointer(mr.Title),
		Description:          helper.CopyPointer(mr.Description),
		SourceBranch:         helper.GetPointer(mr.SourceBranch),
		TargetBranch:         helper.GetPointer(mr.TargetBranch),
		AssigneeIDs:          assigneeIDs,
		ReviewerIDs:          reviewerIDs,
		TargetProjectID:      helper.GetPointer(int(mr.ProjectID)),
		RemoveSourceBranch:   helper.CopyPointer(mr.RemoveSourceBranch),
		Squash:               helper.CopyPointer(mr.Squash),
		ApprovalsBeforeMerge: helper.CopyPointer(mr.ApprovalsBeforeMerge),
	}
}

func GetMRResponse(response *gitlab.MergeRequest) (url string) {
	if response == nil {
		return url
	}
	return response.WebURL
}
