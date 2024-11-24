package config

type Harvest struct {
	AccountID string `mapstructure:"harvest_account_id" validate:"required"`
	Token     string `mapstructure:"harvest_token" validate:"required"`
}
