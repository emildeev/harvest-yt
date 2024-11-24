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
