package devto

import (
	"github.com/olebedev/config"

	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "dev.to | News Feed"
)

// Settings defines the configuration options for this module
type Settings struct {
	common *cfg.Common

	numberOfArticles int    `help:"Number of stories to show. Default is 10" optional:"true"`
	contentTag       string `help:"List articles from a specific tag. Default is empty" optional:"true"`
	contentUsername  string `help:"List articles from a specific user. Default is empty" optional:"true"`
	contentState     string `help:"Order the feed by fresh/rising. Default is rising" optional:"true"`
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, yamlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:           cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, yamlConfig, globalConfig),
		numberOfArticles: yamlConfig.UInt("numberOfArticles", 10),
		contentTag:       yamlConfig.UString("contentTag", ""),
		contentUsername:  yamlConfig.UString("contentUsername", ""),
		contentState:     yamlConfig.UString("contentState", ""),
	}

	return &settings
}
