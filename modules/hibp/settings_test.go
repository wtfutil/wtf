package hibp

import (
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

func newTestSettings(since, apiKey string, accounts []string) *Settings {
	return &Settings{
		Common: &cfg.Common{
			Title: "HIBP",
		},
		apiKey:   apiKey,
		accounts: accounts,
		since:    since,
	}
}

func TestSettings_HasSince(t *testing.T) {
	tests := []struct {
		name     string
		since    string
		expected bool
	}{
		{"empty since", "", false},
		{"valid since", "2019-06-22", true},
		{"malformed since", "not-a-date", false},
		{"partial date", "2019-06", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := newTestSettings(tt.since, "", nil)
			assert.Equal(t, tt.expected, settings.HasSince())
		})
	}
}

func TestSettings_SinceDate(t *testing.T) {
	tests := []struct {
		name        string
		since       string
		expectError bool
		expected    time.Time
	}{
		{
			name:        "valid date",
			since:       "2019-06-22",
			expectError: false,
			expected:    time.Date(2019, 6, 22, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "empty date",
			since:       "",
			expectError: true,
		},
		{
			name:        "malformed date",
			since:       "banana",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := newTestSettings(tt.since, "", nil)
			dt, err := settings.SinceDate()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, tt.expected.Equal(dt))
			}
		})
	}
}

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name                string
		yaml                string
		expectedAccounts    []string
		expectedSince       string
		expectedOkColor     string
		expectedPwnedColor  string
		expectMinRefreshMet bool
	}{
		{
			name: "defaults applied when not configured",
			yaml: `
accounts:
  - test@example.com
`,
			expectedAccounts:    []string{"test@example.com"},
			expectedSince:       "",
			expectedOkColor:     "white",
			expectedPwnedColor:  "red",
			expectMinRefreshMet: true,
		},
		{
			name: "custom colors and since",
			yaml: `
accounts:
  - one@example.com
  - two@example.com
since: 2019-06-22
colors:
  ok: green
  pwned: orange
`,
			expectedAccounts:    []string{"one@example.com", "two@example.com"},
			expectedSince:       "2019-06-22",
			expectedOkColor:     "green",
			expectedPwnedColor:  "orange",
			expectMinRefreshMet: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NoError(t, err)

			globalConfig, err := config.ParseYaml(`global: {}`)
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("hibp", ymlConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.expectedAccounts, settings.accounts)
			assert.Equal(t, tt.expectedSince, settings.since)
			assert.Equal(t, tt.expectedOkColor, settings.ok)
			assert.Equal(t, tt.expectedPwnedColor, settings.pwned)

			if tt.expectMinRefreshMet {
				assert.GreaterOrEqual(t, settings.RefreshInterval, minRefreshInterval)
			}
		})
	}
}

func TestNewSettingsFromYAML_RefreshIntervalFloor(t *testing.T) {
	// A configured refresh interval below the HIBP-enforced minimum should be
	// bumped up to the minimum so we don't hammer the API.
	ymlConfig, err := config.ParseYaml(`
accounts:
  - test@example.com
refreshInterval: 5
`)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml(`global: {}`)
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("hibp", ymlConfig, globalConfig)

	assert.Equal(t, minRefreshInterval, settings.RefreshInterval)
}

func TestConfigText(t *testing.T) {
	widget := &Widget{}
	text := widget.ConfigText()

	assert.NotEmpty(t, text)
}
