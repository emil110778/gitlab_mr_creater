package gitlabcore

type MRID int

type CreatedMRInfo struct {
	URL string
	ID  MRID
	MRInfo
}

type MRInfo struct {
	Title        string
	SourceBranch string
	TargetBranch string
	ProjectID    ProjectID
	MROptionalInfo
}

type FilterMR struct {
	ProjectID    ProjectID
	SourceBranch *string
	TargetBranch *string
	State        *MRState
}

type MRState string

const (
	MRStateOpened MRState = "opened"
	MRStateClosed MRState = "closed"
)

type MROptionalInfo struct {
	Description          *string
	Draft                bool
	Assignees            []UserID
	Reviewers            []UserID
	RemoveSourceBranch   *bool
	Squash               *bool
	ApprovalsBeforeMerge *int
}

type MRUpdateInfo struct {
	ID           MRID
	Title        *string
	TargetBranch *string
	ProjectID    ProjectID
	MROptionalUpdateInfo
}

type MROptionalUpdateInfo struct {
	Description        *string
	Draft              *bool
	Assignees          *[]UserID
	Reviewers          *[]UserID
	RemoveSourceBranch *bool
	Squash             *bool
}

type ResultMRInfo struct {
	URL    string
	Brunch string
	Err    error
}
