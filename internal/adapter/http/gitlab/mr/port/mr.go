package port

import (
	"strings"

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

func UpdateMRRequest(mr gitlabcore.MRUpdateInfo) gitlab.UpdateMergeRequestOptions {
	var (
		assigneeIDs *[]int
		reviewerIDs *[]int
	)

	if mr.Draft != nil && mr.Title != nil && *mr.Draft {
		*mr.Title = draftPrefix + *mr.Title
	}

	if mr.Assignees != nil && len(*mr.Assignees) != 0 {
		assigneeIDs = helper.GetPointer(helper.SliceToInt(*mr.Assignees))
	}

	if mr.Reviewers != nil && len(*mr.Reviewers) != 0 {
		reviewerIDs = helper.GetPointer(helper.SliceToInt(*mr.Reviewers))
	}

	return gitlab.UpdateMergeRequestOptions{
		Title:              helper.CopyPointer(mr.Title),
		Description:        helper.CopyPointer(mr.Description),
		TargetBranch:       helper.CopyPointer(mr.TargetBranch),
		AssigneeIDs:        assigneeIDs,
		ReviewerIDs:        reviewerIDs,
		RemoveSourceBranch: helper.CopyPointer(mr.RemoveSourceBranch),
		Squash:             helper.CopyPointer(mr.Squash),
	}
}

func GetMRResponseURL(response *gitlab.MergeRequest) (url string) {
	if response == nil {
		return url
	}
	return response.WebURL
}

func GetMRResponse(response gitlab.MergeRequest) (mr gitlabcore.CreatedMRInfo) {
	var draft bool

	title := strings.TrimPrefix(response.Title, draftPrefix)
	if response.Title != title {
		draft = true
	}

	ret := gitlabcore.CreatedMRInfo{
		ID:  gitlabcore.MRID(response.IID),
		URL: response.WebURL,
		MRInfo: gitlabcore.MRInfo{
			Title:        title,
			SourceBranch: response.SourceBranch,
			TargetBranch: response.TargetBranch,
			ProjectID:    gitlabcore.ProjectID(response.ProjectID),
			MROptionalInfo: gitlabcore.MROptionalInfo{
				Description:          helper.GetPointer(response.Description),
				Draft:                draft,
				Assignees:            basicUsersToCoreIDs(response.Assignees),
				Reviewers:            basicUsersToCoreIDs(response.Reviewers),
				RemoveSourceBranch:   helper.GetPointer(response.ShouldRemoveSourceBranch),
				Squash:               helper.GetPointer(response.Squash),
				ApprovalsBeforeMerge: helper.GetPointer(response.ApprovalsBeforeMerge),
			},
		},
	}

	return ret
}

func basicUsersToCoreIDs(users []*gitlab.BasicUser) []gitlabcore.UserID {
	if users == nil {
		return nil
	}

	ids := make([]gitlabcore.UserID, 0, len(users))
	for _, user := range users {
		if user == nil {
			continue
		}
		ids = append(ids, basicUserToCoreID(*user))
	}
	return ids
}

func basicUserToCoreID(user gitlab.BasicUser) gitlabcore.UserID {
	return gitlabcore.UserID(user.ID)
}
