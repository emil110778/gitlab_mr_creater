package mr

import (
	"github.com/xanzy/go-gitlab"
)

type clientI interface {
	CreateMergeRequest(
		pid interface{}, opt *gitlab.CreateMergeRequestOptions, options ...gitlab.RequestOptionFunc,
	) (*gitlab.MergeRequest, *gitlab.Response, error)
}

type Adapter struct {
	client clientI
}

func New(client clientI) *Adapter {
	return &Adapter{
		client: client,
	}
}
