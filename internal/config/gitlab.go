package config

type Gitlab struct {
	Host  string `mapstructure:"gitlab_host" validate:"required"`
	Token string `mapstructure:"gitlab_token" validate:"required"`
}
