package config

type Repo struct {
	MainBrunch         string   `mapstructure:"main_brunch" validate:"required"`
	AdditionalBrunches []string `mapstructure:"additional_brunches"`
}
