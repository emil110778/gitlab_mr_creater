package config

type YTracker struct {
	Token string `mapstructure:"y_tracker_token" validate:"required"`
	OrgID int    `mapstructure:"y_tracker_org_id" validate:"required"`
}
