package project

import "github.com/xanzy/go-gitlab"

type clientI interface {
	ListProjects(
		opt *gitlab.ListProjectsOptions, options ...gitlab.RequestOptionFunc,
	) ([]*gitlab.Project, *gitlab.Response, error)
}

type Adapter struct {
	client clientI
}

func New(client clientI) *Adapter {
	return &Adapter{
		client: client,
	}
}
