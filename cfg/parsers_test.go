package cfg

import (
	"testing"

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
