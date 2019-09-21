package azuredevops

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "azuredevops"

// Settings defines the configuration options for this module
type Settings struct {
	common      *cfg.Common
	labelColor  string
	apiToken    string
	orgURL      string
	projectName string
	maxRows     int
}

// NewSettingsFromYAML creates and returns an instance of Settings with configuration options populated
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:      cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),
		labelColor:  ymlConfig.UString("labelColor", "white"),
		apiToken:    ymlConfig.UString("apiToken", "api token not specified"),
		orgURL:      ymlConfig.UString("orgURL", "org url not specified"),
		projectName: ymlConfig.UString("projectName", "project name not specified"),
		maxRows:     ymlConfig.UInt("maxRows", 3),
	}

	return &settings
}
