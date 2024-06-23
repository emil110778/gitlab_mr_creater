package port

import (
	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
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
