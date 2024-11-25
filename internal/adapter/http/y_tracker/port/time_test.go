package port

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_unmarshalDuration(t *testing.T) {
	testCase := []struct {
		name     string
		duration string
		res      time.Duration
	}{
		{
			name:     "hours",
			duration: "PT3H",
			res:      3 * time.Hour,
		},
		{
			name:     "minutes",
			duration: "PT30M",
			res:      30 * time.Minute,
		},
		{
			name:     "Hours and minutes",
			duration: "PT3H30M",
			res:      3*time.Hour + 30*time.Minute,
		},
	}

	for _, tc := range testCase {
		t.Run(
			tc.name, func(t *testing.T) {
				res, err := unmarshalDuration(tc.duration)
				assert.NoError(t, err)
				assert.Equal(t, tc.res, res)
			},
		)
	}
}

func Test_marshalDuration(t *testing.T) {
	testCase := []struct {
		name     string
		duration time.Duration
		res      string
	}{
		{
			name:     "hours",
			duration: 3 * time.Hour,
			res:      "PT3H",
		},
		{
			name:     "minutes",
			duration: 30 * time.Minute,
			res:      "PT30M",
		},
		{
			name:     "Hours and minutes",
			duration: 3*time.Hour + 30*time.Minute,
			res:      "PT3H30M",
		},
	}

	for _, tc := range testCase {
		t.Run(
			tc.name, func(t *testing.T) {
				res := marshalDuration(tc.duration)
				assert.Equal(t, tc.res, res)
			},
		)
	}
}

func Test_unmarshalTime(t *testing.T) {
	testCase := []struct {
		name string
		time string
		res  time.Time
	}{
		{
			name: "hours",
			time: "2021-09-21T15:30:00.000+0500",
			res:  time.Date(2021, 9, 21, 15, 30, 0, 0, time.FixedZone("", 5*60*60)),
		},
	}

	for _, tc := range testCase {
		t.Run(
			tc.name, func(t *testing.T) {
				res, err := unmarshalTime(tc.time)
				assert.NoError(t, err)
				assert.Equal(t, tc.res, res)
			},
		)
	}
}

func Test_marshalTime(t *testing.T) {
	testCase := []struct {
		name string
		time time.Time
		res  string
	}{
		{
			name: "hours",
			time: time.Date(2021, 9, 21, 15, 30, 0, 0, time.FixedZone("", 5*60*60)),
			res:  "2021-09-21T15:30:00+0500",
		},
	}

	for _, tc := range testCase {
		t.Run(
			tc.name, func(t *testing.T) {
				res := marshalTime(tc.time)
				assert.Equal(t, tc.res, res)
			},
		)
	}
}
