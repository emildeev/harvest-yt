package port

import (
	"time"

	"github.com/sosodev/duration"
)

func unmarshalTime(duration string) (time.Time, error) {
	return time.Parse(time.RFC3339, duration)
}

func marshalTime(t time.Time) string {
	return t.Format(time.RFC3339)
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
