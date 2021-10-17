package cfg

import (
	"fmt"
	"time"

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

// ParseTimeString takes a configuration key and attempts to parse it first as an int
// and then as a duration (int + time unit)
func ParseTimeString(cfg *config.Config, configKey string, defaultValue string) time.Duration {
	i, err := cfg.Int(configKey)
	if err == nil {
		return time.Duration(i) * time.Second
	}

	str := cfg.UString(configKey, defaultValue)
	d, err := time.ParseDuration(str)
	if err == nil {
		return d
	}

	return time.Second
}
