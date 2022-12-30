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

type colors struct {
	source      string `help:"Color to use for feed source titles." optional:"true" default:"green"`
	publishDate string `help:"Color to use for publish dates." optional:"true" default:"orange"`
}

// auth stores [username, password]-credentials for private RSS feeds using Basic Auth
type auth struct {
	username string
	password string
}

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	colors

	feeds           []string        `help:"An array of RSS and Atom feed URLs"`
	feedLimit       int             `help:"The maximum number of stories to display for each feed"`
	showSource      bool            `help:"Wether or not to show feed source in front of item titles." values:"true or false" optional:"true" default:"true"`
	showPublishDate bool            `help:"Wether or not to show publish date in front of item titles." values:"true or false" optional:"true" default:"false"`
	dateFormat      string          `help:"Date format to use for publish dates" values:"Any valid Go time layout which is handled by Time.Format" optional:"true" default:"Jan 02"`
	credentials     map[string]auth `help:"Map of private feed URLs with required authentication credentials"`
	disableHTTP2    bool            `help:"Wether or not to use the HTTP/2 protocol. Certain sites, such as reddit.com, will not work unless HTTP/2 is disabled." values:"true or false" optional:"true" default:"false"`
	userAgent       string          `help:"HTTP User-Agent to use when fetching RSS feeds." optional:"true"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig, globalConfig *config.Config) *Settings {
	settings := &Settings{
		Common:          cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		feeds:           utils.ToStrs(ymlConfig.UList("feeds")),
		feedLimit:       ymlConfig.UInt("feedLimit", -1),
		showSource:      ymlConfig.UBool("showSource", true),
		showPublishDate: ymlConfig.UBool("showPublishDate", false),
		dateFormat:      ymlConfig.UString("dateFormat", "Jan 02"),
		credentials:     make(map[string]auth),
		disableHTTP2:    ymlConfig.UBool("disableHTTP2", false),
		userAgent:       ymlConfig.UString("userAgent", "wtfutil (https://github.com/wtfutil/wtf)"),
	}

	settings.colors.source = ymlConfig.UString("colors.source", "green")
	settings.colors.publishDate = ymlConfig.UString("colors.publishDate", "orange")

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
