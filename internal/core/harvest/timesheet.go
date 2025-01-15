package harvestcore

import "time"

type TimeEntries []TimeEntry

type TimeEntry struct {
	ID           int64
	Client       string
	Project      string
	Task         string
	Hours        time.Duration
	RoundedHours time.Duration
	Notes        string
	IsRunning    bool
	CreatedAt    time.Time
}
