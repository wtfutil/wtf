package app

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/modules/clocks"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MakeWidget(nil, nil, tt.moduleName, tt.config, make(chan bool))
			assert.IsType(t, tt.expected, actual)
		})
	}
}
