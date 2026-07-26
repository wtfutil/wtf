package hibp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBreach_BreachDate(t *testing.T) {
	tests := []struct {
		name        string
		date        string
		expectError bool
		expected    time.Time
	}{
		{
			name:        "valid date",
			date:        "2013-10-04",
			expectError: false,
			expected:    time.Date(2013, 10, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "another valid date",
			date:        "2019-06-22",
			expectError: false,
			expected:    time.Date(2019, 6, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "empty date",
			date:        "",
			expectError: true,
		},
		{
			name:        "malformed date",
			date:        "not-a-date",
			expectError: true,
		},
		{
			name:        "wrong format",
			date:        "10/04/2013",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			br := &Breach{Date: tt.date}
			dt, err := br.BreachDate()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, tt.expected.Equal(dt))
			}
		})
	}
}
