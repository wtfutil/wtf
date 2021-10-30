package cfg

import (
	"testing"
	"time"

	"github.com/olebedev/config"
)

func Test_ParseAsMapOrList(t *testing.T) {
	tests := []struct {
		name          string
		configKey     string
		yaml          string
		expectedCount int
	}{
		{
			name:          "as empty set",
			configKey:     "data",
			yaml:          "",
			expectedCount: 0,
		},
		{
			name:          "as map",
			configKey:     "data",
			yaml:          "data:\n  a: cat\n  b: dog",
			expectedCount: 2,
		},
		{
			name:          "as list",
			configKey:     "data",
			yaml:          "data:\n  - cat\n  - dog\n  - rat\n",
			expectedCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			if err != nil {
				t.Errorf("\nexpected: no error\n     got: %v", err)
			}

			actual := ParseAsMapOrList(ymlConfig, tt.configKey)

			if tt.expectedCount != len(actual) {
				t.Errorf("\nexpected: %d\n     got: %d", tt.expectedCount, len(actual))
			}
		})
	}
}

func Test_ParseTimeString(t *testing.T) {
	tests := []struct {
		name          string
		configKey     string
		yaml          string
		expectedCount time.Duration
	}{
		{
			name:          "normal integer",
			configKey:     "refreshInterval",
			yaml:          "refreshInterval: 3",
			expectedCount: 3 * time.Second,
		},
		{
			name:          "microseconds",
			configKey:     "refreshInterval",
			yaml:          "refreshInterval: 5µs",
			expectedCount: 5 * time.Microsecond,
		},
		{
			name:          "microseconds different notation",
			configKey:     "refreshInterval",
			yaml:          "refreshInterval: 5us",
			expectedCount: 5 * time.Microsecond,
		},
		{
			name:          "mixed duration",
			configKey:     "refreshInterval",
			yaml:          "refreshInterval: 2h45m",
			expectedCount: 2*time.Hour + 45*time.Minute,
		},
		{
			name:          "default",
			configKey:     "refreshInterval",
			yaml:          "",
			expectedCount: 60 * time.Second,
		},
		{
			name:          "bad input",
			configKey:     "refreshInterval",
			yaml:          "refreshInterval: abc",
			expectedCount: 1 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			if err != nil {
				t.Errorf("\nexpected: no error\n     got: %v", err)
			}

			actual := ParseTimeString(ymlConfig, tt.configKey, "60s")

			if tt.expectedCount != actual {
				t.Errorf("\nexpected: %d\n     got: %v", tt.expectedCount, actual)
			}
		})
	}
}
