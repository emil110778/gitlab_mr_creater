package gitlab

import (
	"context"

	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
)

type mrAdapterI interface {
	CreateMR(ctx context.Context, mr gitlabcore.MRInfo) (url string, err error)
}
type projectAdapterI interface {
	GetProjects(ctx context.Context) (gitlabProjects []gitlabcore.Project, err error)
}

type Service struct {
	mrAdapter      mrAdapterI
	projectAdapter projectAdapterI
}

func New(adapter mrAdapterI, projectAdapter projectAdapterI) *Service {
	return &Service{
		mrAdapter:      adapter,
		projectAdapter: projectAdapter,
	}
}
