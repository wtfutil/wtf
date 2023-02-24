package app

import (
	"fmt"
	"testing"

	"github.com/logrusorgru/aurora/v4"
	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/wtf"
)

const (
	valid = `
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

	invalid = `
wtf:
  mods:
    clocks:
      enabled: true
      position:
        top: abc
        left: 0
        height: 1
        width: 1
      refreshInterval: 30`
)

func Test_NewModuleValidator(t *testing.T) {
	assert.IsType(t, &ModuleValidator{}, NewModuleValidator())
}

func Test_validate(t *testing.T) {
	tests := []struct {
		name       string
		moduleName string
		config     *config.Config
		expected   []string
	}{
		{
			name:       "valid config",
			moduleName: "clocks",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(valid)
				return cfg
			}(),
			expected: []string{},
		},
		{
			name:       "invalid config",
			moduleName: "clocks",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(invalid)
				return cfg
			}(),
			expected: []string{
				fmt.Sprintf("%s in %s configuration", aurora.Red("Errors"), aurora.Yellow("clocks.position")),
				fmt.Sprintf(
					" - Invalid value for %s:	0	%s strconv.ParseInt: parsing \"abc\": invalid syntax",
					aurora.Yellow("top"),
					aurora.Red("Error:"),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := MakeWidget(nil, nil, tt.moduleName, tt.config, make(chan bool))

			if widget == nil {
				t.Logf("Failed to create widget %s", tt.moduleName)
				t.FailNow()
			}

			errs := validate([]wtf.Wtfable{widget})

			if len(tt.expected) == 0 {
				assert.Empty(t, errs)
			} else {
				assert.NotEmpty(t, errs)

				var actual []string
				for _, err := range errs {
					actual = append(actual, err.errorMessages()...)
				}

				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}
