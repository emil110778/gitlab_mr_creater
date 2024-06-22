package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP HTTP `mapstructure:"http" validate:"required"`
	Repo Repo `mapstructure:"repo" validate:"required"`
}

func New() (Config, error) {
	errWrap := func(err error) error {
		return fmt.Errorf("error load config: %w", err)
	}
	config := Config{}

	err := viper.Unmarshal(&config)
	if err != nil {
		return Config{}, errWrap(err)
	}

	validate := validator.New()
	if err = validate.Struct(&config); err != nil {
		return Config{}, errWrap(err)
	}

	return config, nil
}

func NewWithoutValidate() Config {
	config := Config{}
	_ = viper.Unmarshal(&config)
	return config
}
