package projecttemplate

import "github.com/xanzy/go-gitlab"

type clientI interface {
	GetProjectTemplate(
		pid interface{}, templateType string, templateName string, options ...gitlab.RequestOptionFunc,
	) (*gitlab.ProjectTemplate, *gitlab.Response, error)
}

type Adapter struct {
	client clientI
}

func New(client clientI) *Adapter {
	return &Adapter{
		client: client,
	}
}
