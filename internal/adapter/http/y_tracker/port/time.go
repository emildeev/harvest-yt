package port

import (
	"time"

	"github.com/sosodev/duration"
)

const (
	timeFormat = "2006-01-02T15:04:05-0700"
)

func unmarshalTime(duration string) (time.Time, error) {
	return time.Parse(timeFormat, duration)
}

func marshalTime(t time.Time) string {
	return t.Format(timeFormat)
}

func unmarshalDuration(durationStr string) (time.Duration, error) {
	d, err := duration.Parse(durationStr)
	if err != nil {
		return 0, err
	}
	return d.ToTimeDuration(), nil
}

func marshalDuration(d time.Duration) string {
	return duration.FromTimeDuration(d).String()
}
