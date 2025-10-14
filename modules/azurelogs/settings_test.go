package azurelogs

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"

	"github.com/wtfutil/wtf/cfg"
)

func TestSettings_Structure(t *testing.T) {
	// Test Settings structure
	settings := &Settings{
		Common: &cfg.Common{
			Title: "Test Azure Logs",
		},
		Queryfile: "/path/to/query.yml",
	}

	assert.NotNil(t, settings.Common)
	assert.Equal(t, "Test Azure Logs", settings.Title)
	assert.Equal(t, "/path/to/query.yml", settings.Queryfile)
}

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name          string
		configData    map[string]interface{}
		expectedTitle string
		expectedQuery string
	}{
		{
			name: "with custom query file",
			configData: map[string]interface{}{
				"queryFile": "/custom/path/query.yml",
				"title":     "Custom Azure Logs",
			},
			expectedTitle: "Custom Azure Logs",
			expectedQuery: "/custom/path/query.yml",
		},
		{
			name:       "with default values",
			configData: map[string]interface{}{
				// No queryFile specified, should use default empty string
			},
			expectedTitle: defaultTitle, // Should use default title
			expectedQuery: "",           // Should use default empty string
		},
		{
			name: "with empty query file",
			configData: map[string]interface{}{
				"queryFile": "",
				"title":     "Empty Query Azure Logs",
			},
			expectedTitle: "Empty Query Azure Logs",
			expectedQuery: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create YAML config from test data
			ymlConfig, err := config.ParseYaml(yamlFromMap(tt.configData))
			assert.NoError(t, err)

			// Create global config (can be minimal for this test)
			globalConfig, err := config.ParseYaml("global: {}")
			assert.NoError(t, err)

			settings := NewSettingsFromYAML("test-widget", ymlConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.expectedTitle, settings.Title)
			assert.Equal(t, tt.expectedQuery, settings.Queryfile)
		})
	}
}

func TestDefaultConstants(t *testing.T) {
	// Test that default constants are correctly defined
	assert.True(t, defaultFocusable)
	assert.Equal(t, "Azure Logs", defaultTitle)
}

func TestSettings_QueryfileField(t *testing.T) {
	// Test that Queryfile field can be set and retrieved
	settings := &Settings{}

	// Test setting various query file paths
	testPaths := []string{
		"/absolute/path/query.yml",
		"relative/path/query.yml",
		"./current/dir/query.yml",
		"../parent/dir/query.yml",
		"",
	}

	for _, path := range testPaths {
		settings.Queryfile = path
		assert.Equal(t, path, settings.Queryfile)
	}
}

func TestNewSettingsFromYAML_Integration(t *testing.T) {
	// Test with a more complete YAML configuration
	configData := map[string]interface{}{
		"queryFile": "/etc/wtf/azure-query.yml",
		"title":     "Production Azure Logs",
		"enabled":   true,
		"position": map[string]interface{}{
			"top":    0,
			"left":   0,
			"width":  2,
			"height": 1,
		},
		"refreshInterval": "5m",
	}

	ymlConfig, err := config.ParseYaml(yamlFromMap(configData))
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml(`
wtf:
  term: "xterm-256color"
  grid:
    columns: [40, 40, 40]
    rows: [13, 13, 4]
`)
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("azure-logs", ymlConfig, globalConfig)

	assert.NotNil(t, settings)
	assert.Equal(t, "Production Azure Logs", settings.Title)
	assert.Equal(t, "/etc/wtf/azure-query.yml", settings.Queryfile)
	assert.NotNil(t, settings.Common)
}

// Helper function to convert map to YAML string for testing
func yamlFromMap(data map[string]interface{}) string {
	if len(data) == 0 {
		return "{}"
	}

	yaml := ""
	for key, value := range data {
		switch v := value.(type) {
		case string:
			yaml += key + ": \"" + v + "\"\n"
		case bool:
			if v {
				yaml += key + ": true\n"
			} else {
				yaml += key + ": false\n"
			}
		case map[string]interface{}:
			yaml += key + ":\n"
			for subKey, subValue := range v {
				yaml += "  " + subKey + ": " + interfaceToString(subValue) + "\n"
			}
		default:
			yaml += key + ": " + interfaceToString(v) + "\n"
		}
	}
	return yaml
}

// Helper function to convert interface{} to string for YAML
func interfaceToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return "\"" + val + "\""
	case int:
		return string(rune(val + '0'))
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		return "null"
	}
}
