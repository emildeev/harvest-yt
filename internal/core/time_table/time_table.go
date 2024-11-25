package timetablecore

import "time"

type Table []Element

func (t Table) HasErr() bool {
	for _, e := range t {
		if e.Err != nil {
			return true
		}
	}
	return false
}

type Element struct {
	TaskKey   string
	TaskTitle string
	Comment   string
	Duration  time.Duration
	Err       error
	StartTime time.Time
	TimerID   int64
}
