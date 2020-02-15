package cdsfavorites

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "CDS Favorites"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common   *cfg.Common
	token    string `help:"Your CDS API token."`
	apiURL   string `help:"Your CDS API URL."`
	uiURL    string
	hideTags []string `help:"Hide some workflow tags."`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:   cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		token:    ymlConfig.UString("token", ymlConfig.UString("token", os.Getenv("CDS_TOKEN"))),
		apiURL:   ymlConfig.UString("apiURL", os.Getenv("CDS_API_URL")),
		hideTags: utils.ToStrs(ymlConfig.UList("hideTags")),
	}
	return &settings
}
