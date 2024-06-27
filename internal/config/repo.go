package config

type Repo struct {
	MainBranch         string   `mapstructure:"main_branch" validate:"required"`
	AdditionalBranches []string `mapstructure:"additional_branches"`
}
