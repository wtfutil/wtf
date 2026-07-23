package ping

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

func Test_buildhosts(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		expected []Host
	}{
		{
			name: "single host with label",
			yaml: `
hosts:
  - hostname: "example.com"
    label: "Example"
`,
			expected: []Host{
				{Label: "Example", Hostname: "example.com", Up: false},
			},
		},
		{
			name: "single host without label defaults to hostname",
			yaml: `
hosts:
  - hostname: "example.com"
`,
			expected: []Host{
				{Label: "example.com", Hostname: "example.com", Up: false},
			},
		},
		{
			name: "multiple hosts",
			yaml: `
hosts:
  - hostname: "example.com"
    label: "Example"
  - hostname: "google.com"
`,
			expected: []Host{
				{Label: "Example", Hostname: "example.com", Up: false},
				{Label: "google.com", Hostname: "google.com", Up: false},
			},
		},
		{
			name: "host missing hostname is skipped",
			yaml: `
hosts:
  - label: "No hostname"
`,
			expected: []Host{},
		},
		{
			name: "host with empty hostname is skipped",
			yaml: `
hosts:
  - hostname: ""
    label: "Empty hostname"
`,
			expected: []Host{},
		},
		{
			name: "host with numeric label is stringified",
			yaml: `
hosts:
  - hostname: "example.com"
    label: 123
`,
			expected: []Host{
				{Label: "123", Hostname: "example.com", Up: false},
			},
		},
		{
			name:     "no hosts configured",
			yaml:     `hosts: []`,
			expected: []Host{},
		},
		{
			name:     "hosts key missing entirely",
			yaml:     `other: value`,
			expected: []Host{},
		},
		{
			name: "non-map entry in hosts list is skipped",
			yaml: `
hosts:
  - "just a string"
  - hostname: "valid.com"
`,
			expected: []Host{
				{Label: "valid.com", Hostname: "valid.com", Up: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			assert.NoError(t, err)

			hosts := buildhosts(ymlConfig)

			assert.Equal(t, tt.expected, hosts)
		})
	}
}

func Test_NewSettingsFromYAML(t *testing.T) {
	ymlConfig, err := config.ParseYaml(`
hosts:
  - hostname: "example.com"
    label: "Example"
`)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml(`wtf: {}`)
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("ping", ymlConfig, globalConfig)

	assert.NotNil(t, settings)
	assert.NotNil(t, settings.common)
	assert.Equal(t, defaultTitle, settings.common.Title)
	assert.Len(t, settings.hosts, 1)
	assert.Equal(t, "Example", settings.hosts[0].Label)
	assert.Equal(t, "example.com", settings.hosts[0].Hostname)
}

func Test_NewSettingsFromYAML_NoHosts(t *testing.T) {
	ymlConfig, err := config.ParseYaml(`other: value`)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml(`wtf: {}`)
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("ping", ymlConfig, globalConfig)

	assert.NotNil(t, settings)
	assert.Empty(t, settings.hosts)
}

func Test_ConfigText(t *testing.T) {
	widget := &Widget{}
	text := widget.ConfigText()

	// Settings only has unexported fields (common, hosts), so
	// utils.HelpFromInterface has no exported/tagged fields to describe.
	assert.Equal(t, "", text)
}

func Test_defaultConstants(t *testing.T) {
	assert.False(t, defaultFocusable)
	assert.Equal(t, "Pings", defaultTitle)
}
