package cfg

import (
	"fmt"

	"github.com/olebedev/config"
)

// ParseAsMapOrList takes a configuration key and attempts to parse it first as a map
// and then as a list. Map entries are concatenated as "key/value"
func ParseAsMapOrList(ymlConfig *config.Config, configKey string) []string {
	result := []string{}

	mapItems, err := ymlConfig.Map(configKey)
	if err == nil {
		for key, value := range mapItems {
			result = append(result, fmt.Sprintf("%s/%s", value, key))
		}
		return result
	}

	listItems := ymlConfig.UList(configKey)
	for _, listItem := range listItems {
		result = append(result, listItem.(string))
	}

	return result
}
