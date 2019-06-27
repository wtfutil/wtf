package transmission

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	host     string
	https    bool
	password string
	port     int
	url      string
	username string
}

const (
	defaultTitle = "Transmission"
)

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		host:     ymlConfig.UString("host"),
		https:    ymlConfig.UBool("https", false),
		password: ymlConfig.UString("password"),
		port:     ymlConfig.UInt("port", 9091),
		url:      ymlConfig.UString("url", "/transmission/"),
		username: ymlConfig.UString("username", ""),
	}

	return &settings
}
