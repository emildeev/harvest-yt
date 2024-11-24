package ytrackercore

import "time"

type WorkLogs []WorkLog

type WorkLog struct {
	ID        int           `json:"id"`
	TicketKey string        `json:"ticket_key"`
	Comment   string        `json:"comment"`
	Start     time.Time     `json:"start"`
	Duration  time.Duration `json:"duration"`
}
