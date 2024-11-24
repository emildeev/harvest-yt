package config

type Tasks struct {
	SkippedTasks       []string          `mapstructure:"skipped_tasks"`
	DevelopTasks       []string          `mapstructure:"main_branch"`
	CommunicationTasks map[string]string `mapstructure:"additional_branches"`
	StartTime          string            `mapstructure:"start_time" validate:"required"`
}
