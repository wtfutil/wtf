package app

import (
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/modules/clocks"
	"github.com/wtfutil/wtf/modules/system"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

const (
	disabled = `
wtf:
  mods:
    clocks:
      enabled: false
      position:
        top: 0
        left: 0
        height: 1
        width: 1
      refreshInterval: 30`

	enabled = `
wtf:
  mods:
    clocks:
      enabled: true
      position:
        top: 0
        left: 0
        height: 1
        width: 1
      refreshInterval: 30`

	systemEnabled = `
wtf:
  mods:
    system:
      enabled: true
      position:
        top: 0
        left: 0
        height: 1
        width: 1
      refreshInterval: 30`
)

func Test_MakeWidget(t *testing.T) {
	tests := []struct {
		name       string
		moduleName string
		config     *config.Config
		expected   wtf.Wtfable
	}{
		{
			name:       "invalid module",
			moduleName: "",
			config:     &config.Config{},
			expected:   nil,
		},
		{
			name:       "valid disabled module",
			moduleName: "clocks",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(disabled)
				return cfg
			}(),
			expected: nil,
		},
		{
			name:       "valid enabled module",
			moduleName: "clocks",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(enabled)
				return cfg
			}(),
			expected: &clocks.Widget{},
		},
		{
			name:       "valid enabled system module",
			moduleName: "system",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(systemEnabled)
				return cfg
			}(),
			expected: &system.Widget{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MakeWidget(nil, nil, tt.moduleName, tt.config, make(chan bool))
			assert.IsType(t, tt.expected, actual)
		})
	}
}

func Test_buildVersion(t *testing.T) {
	actual := buildVersion()
	assert.NotEmpty(t, actual)
}

func Test_buildDate(t *testing.T) {
	actual := buildDate()

	_, err := time.Parse(utils.TimestampFormat, actual)
	assert.NoError(t, err)
}
