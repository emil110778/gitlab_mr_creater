package gitlab

import (
	"context"
	"fmt"

	gitlabcore "github.com/emil110778/gitlab_mr_creator/internal/core/gitlab"
)

func (service *Service) CreateMR(ctx context.Context, mr gitlabcore.MRInfo) (url string, err error) {
	mrID, err := service.mrAdapter.CreateMR(ctx, mr)
	if err != nil {
		return url, fmt.Errorf("CreateMR: %w", err)
	}
	return mrID, nil
}
