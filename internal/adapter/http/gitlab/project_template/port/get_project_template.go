package port

import (
	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
	"github.com/xanzy/go-gitlab"
)

func GetGetMRTemplateResponse(response *gitlab.ProjectTemplate) (mrTemplate gitlabcore.MRTemplate) {
	if response == nil {
		return mrTemplate
	}

	return gitlabcore.MRTemplate{
		Description: response.Content,
	}
}
