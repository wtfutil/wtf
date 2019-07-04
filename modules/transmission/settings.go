package transmission

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	host     string `help:"The address of the machine the Transmission daemon is running on"`
	https    bool   `help:"Whether or not to connect to the host via HTTPS"`
	password string `help:"The password for the transmission user"`
	port     int    `help:"The port to connect to the Transmission daemon on"`
	url      string `help:"The RPC URI that the daemon is accessible at"`
	username string `help:"The username for the transmission user"`
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
