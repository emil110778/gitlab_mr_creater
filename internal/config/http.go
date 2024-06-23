package config

type HTTP struct {
	Gitlab   Gitlab   `mapstructure:"gitlab" validate:"required"`
	YTracker YTracker `mapstructure:"y_tracker" validate:"required"`
}
