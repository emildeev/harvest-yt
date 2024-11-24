package config

type HTTP struct {
	Harvest  Harvest  `mapstructure:"hytmigrator" validate:"required"`
	YTracker YTracker `mapstructure:"y_tracker" validate:"required"`
}
