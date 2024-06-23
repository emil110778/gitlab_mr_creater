package gitlabcore

type MRID int

type MRInfo struct {
	Title        string
	SourceBranch string
	TargetBranch string
	ProjectID    ProjectID
	MROptionalInfo
}

type MROptionalInfo struct {
	Description          *string
	Draft                bool
	Assignees            []UserID
	Reviewers            []UserID
	RemoveSourceBranch   *bool
	Squash               *bool
	ApprovalsBeforeMerge *int
}

type CreatedMRInfo struct {
	URL    string
	Brunch string
	Err    error
}
