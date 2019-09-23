package subreddit

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
)

// Settings contains the settings for the subreddit view
type Settings struct {
	common *cfg.Common

	subreddit     string `help:"Subreddit to look at" optional:"false"`
	numberOfPosts int    `help:"Number of posts to show. Default is 10." optional:"true"`
	sortOrder     string `help:"Sort order for the posts (hot, new, rising, top), default hot" optional:"true"`
}

// NewSettingsFromYAML creates the settings for this module from a yaml file
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	subreddit := ymlConfig.UString("subreddit")
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, subreddit, defaultFocusable, ymlConfig, globalConfig),

		numberOfPosts: ymlConfig.UInt("numberOfPosts", 10),
		sortOrder:     ymlConfig.UString("sortOrder", "hot"),
		subreddit:     subreddit,
	}

	return &settings
}
