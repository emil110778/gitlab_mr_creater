package config

type HTTP struct {
	Gitlab Gitlab `mapstructure:"gitlab" validate:"required"`
}
