package hibp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStatus(t *testing.T) {
	breaches := []Breach{{Name: "Adobe", Date: "2013-10-04"}}

	stat := NewStatus("user@example.com", breaches)

	assert.NotNil(t, stat)
	assert.Equal(t, "user@example.com", stat.Account)
	assert.Equal(t, breaches, stat.Breaches)
}

func TestStatus_HasBeenCompromised(t *testing.T) {
	tests := []struct {
		name     string
		stat     *Status
		expected bool
	}{
		{"nil status", nil, false},
		{"no breaches", NewStatus("acct", []Breach{}), false},
		{"nil breaches slice", NewStatus("acct", nil), false},
		{"one breach", NewStatus("acct", []Breach{{Name: "Adobe"}}), true},
		{"multiple breaches", NewStatus("acct", []Breach{{Name: "Adobe"}, {Name: "LinkedIn"}}), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.stat.HasBeenCompromised())
		})
	}
}

func TestStatus_Len(t *testing.T) {
	tests := []struct {
		name     string
		stat     *Status
		expected int
	}{
		{"nil status", nil, 0},
		{"nil breaches", NewStatus("acct", nil), 0},
		{"empty breaches", NewStatus("acct", []Breach{}), 0},
		{"single breach", NewStatus("acct", []Breach{{Name: "Adobe"}}), 1},
		{"two breaches", NewStatus("acct", []Breach{{Name: "Adobe"}, {Name: "LinkedIn"}}), 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.stat.Len())
		})
	}
}
