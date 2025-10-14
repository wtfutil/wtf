package azurelogs

import (
	_ "embed"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// QueryFile represents the structure of a query configuration file
type QueryFile struct {
	Title          string   `yaml:"title"`                 // Display title for the query
	SubscriptionID string   `yaml:"azure_subscription_id"` // Azure subscription ID
	WorkspaceID    string   `yaml:"azure_workspace_id"`    // Log Analytics workspace ID
	Columns        []string `yaml:"columns"`               // Expected column names
	Query          string   `yaml:"query"`                 // KQL query string
}

// readQueryFile reads and parses a query configuration file
func readQueryFile(sess *Session, queryPath string) error {
	file, err := os.OpenFile(queryPath, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}

	filename := file.Name()
	if len(filename) > 5 && filename[len(filename)-5:] == ".yaml" {
		var configFile QueryFile
		configFile, err = readQueryFileContent(queryPath)
		if err != nil {
			return err
		}

		sess.QueryFile = configFile
	} else {
		return fmt.Errorf("invalid query file format: %s, expected .yaml", filename)
	}

	return nil
}

// readQueryFileContent reads a single config file and returns a QueryFile struct
func readQueryFileContent(filePath string) (QueryFile, error) {
	var configFile QueryFile
	data, err := os.ReadFile(filePath)
	if err != nil {
		return configFile, fmt.Errorf("failed to read query config file %s: %w", filePath, err)
	}

	err = yaml.Unmarshal(data, &configFile)
	if err != nil {
		return configFile, fmt.Errorf("failed to parse YAML in config file %s: %w", filePath, err)
	}

	return configFile, nil
}
