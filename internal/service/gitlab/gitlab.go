package gitlab

import (
	"context"
	"regexp"

	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
)

type mrAdapterI interface {
	CreateMR(ctx context.Context, mr gitlabcore.MRInfo) (url string, err error)
}
type projectAdapterI interface {
	GetProjects(ctx context.Context) (gitlabProjects []gitlabcore.Project, err error)
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
