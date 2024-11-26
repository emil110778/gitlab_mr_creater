package gitlab

import (
	"context"
	"regexp"

	gitlabcore "github.com/emildeev/gitlab_helper/internal/core/gitlab"
)

type mrAdapterI interface {
	Create(ctx context.Context, mr gitlabcore.MRInfo) (url string, err error)
	List(ctx context.Context, filter gitlabcore.FilterMR) (mrs []gitlabcore.CreatedMRInfo, err error)
	Update(ctx context.Context, mr gitlabcore.MRUpdateInfo) (url string, err error)
}
type projectAdapterI interface {
	GetProjectsByProjectName(ctx context.Context, projectName *string) (gitlabProjects []gitlabcore.Project, err error)
}
type projectTemplateAdapterI interface {
	GetMRTemplate(ctx context.Context, projectID int, templateName string) (mrTemplate gitlabcore.MRTemplate, err error)
}

type Service struct {
	mrAdapter              mrAdapterI
	projectAdapter         projectAdapterI
	projectTemplateAdapter projectTemplateAdapterI

	mrTemplateFieldRegExp *regexp.Regexp
}

func New(
	adapter mrAdapterI, projectAdapter projectAdapterI, projectTemplateAdapter projectTemplateAdapterI,
) (*Service, error) {
	mrTemplateFieldRegExp, err := regexp.Compile("[*]{2}.+:[*]{2}[[:blank:]]*X{3}")
	if err != nil {
		return nil, err
	}

	return &Service{
		mrAdapter:              adapter,
		projectAdapter:         projectAdapter,
		projectTemplateAdapter: projectTemplateAdapter,
		mrTemplateFieldRegExp:  mrTemplateFieldRegExp,
	}, nil
}
