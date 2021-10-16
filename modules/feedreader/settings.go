package feedreader

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Feed Reader"
)

// Auth stores [username, password]-credentials for Basic Auth
type Auth struct {
	username string
	password string
}

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	feeds       []string `help:"An array of RSS and Atom feed URLs"`
	feedLimit   int      `help:"The maximum number of stories to display for each feed"`
	credentials []Auth   `help:"List of [user, password]-pairs to try basic authentication on a feed"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		feeds:       utils.ToStrs(ymlConfig.UList("feeds")),
		feedLimit:   ymlConfig.UInt("feedLimit", -1),
		credentials: listifyCredentials(ymlConfig, globalConfig),
	}

	return settings
}

// listifyCredentials converts a list of [user, password]-pairs to a slice auth structs
func listifyCredentials(ymlConfig *config.Config, globalConfig *config.Config) []Auth {
	list, err := ymlConfig.List("auth")
	if err != nil {
		return []Auth{}
	}

	result := make([]Auth, 0)
	for _, entry := range list {
		credentials, ok := entry.([]interface{})
		if !ok || len(credentials) != 2 {
			continue
		}

		user, ok1 := credentials[0].(string)
		pass, ok2 := credentials[1].(string)
		if !ok1 || !ok2 {
			continue
		}

		result = append(result, Auth{username: user, password: pass})
	}
	return result
}
