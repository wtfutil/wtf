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

// auth stores [username, password]-credentials for private RSS feeds using Basic Auth
type auth struct {
	username string
	password string
}

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	feeds        []string        `help:"An array of RSS and Atom feed URLs"`
	feedLimit    int             `help:"The maximum number of stories to display for each feed"`
	credentials  map[string]auth `help:"Map of private feed URLs with required authentication credentials"`
	disableHTTP2 bool            `help:"Wether or not to use the HTTP/2 protocol. Certain sites, such as reddit.com, will not work unless HTTP/2 is disabled." values:"true or false" optional:"true" default:"false"`
	userAgent    string          `help:"HTTP User-Agent to use when fetching RSS feeds." optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common:       cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		feeds:        utils.ToStrs(ymlConfig.UList("feeds")),
		feedLimit:    ymlConfig.UInt("feedLimit", -1),
		credentials:  make(map[string]auth),
		disableHTTP2: ymlConfig.UBool("disableHTTP2", false),
		userAgent:    ymlConfig.UString("userAgent", "wtfutil (https://github.com/wtfutil/wtf)"),
	}

	// If feeds cannot be parsed as list try parsing as a map with username+password fields
	if len(settings.feeds) == 0 {
		credentials := make(map[string]auth)
		feeds := make([]string, 0)
		for url, creds := range ymlConfig.UMap("feeds") {
			parsed, ok := creds.(map[string]interface{})
			if !ok {
				continue
			}

			user, ok := parsed["username"].(string)
			if !ok {
				continue
			}
			pass, ok := parsed["password"].(string)
			if !ok {
				continue
			}

			credentials[url] = auth{
				username: user,
				password: pass,
			}
			feeds = append(feeds, url)
		}
		settings.feeds = feeds
		settings.credentials = credentials
	}

	return settings
}
