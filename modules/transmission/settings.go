package transmission

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Transmission"
)

// Settings defines the configuration properties for this module
type Settings struct {
	*cfg.Common

	host         string `help:"The address of the machine the Transmission daemon is running on"`
	https        bool   `help:"Whether or not to connect to the host via HTTPS"`
	password     string `help:"The password for the Transmission user"`
	port         uint16 `help:"The port to connect to the Transmission daemon on"`
	url          string `help:"The RPC URI that the daemon is accessible at"`
	username     string `help:"The username of the Transmission user"`
	hideComplete bool   `help:"Hide the torrents that are finished downloading"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		host:         ymlConfig.UString("host"),
		https:        ymlConfig.UBool("https", false),
		password:     ymlConfig.UString("password"),
		port:         uint16(ymlConfig.UInt("port", 9091)),
		url:          ymlConfig.UString("url", ""),
		username:     ymlConfig.UString("username", ""),
		hideComplete: ymlConfig.UBool("hideComplete", false),
	}

	return &settings
}
