package mail

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "mail"
)

// Settings defines the configuration properties for this module
type Settings struct {
	common *cfg.Common

	imapAddress     string `help:"The address and port to the IMAP server" values:"imap.example.com:993"`
	username        string `help:"The username to log into the email account"`
	password        string `help:"The password of the email account"`
	defaultPageSize int    `help:"The default number of messages to display per page" values:"Numbers greater than 0"`
}

// NewSettingsFromYAML creates a new settings instance from a YAML config block
func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:          cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		imapAddress:     ymlConfig.UString("imapAddress"),
		username:        ymlConfig.UString("username"),
		password:        ymlConfig.UString("password"),
		defaultPageSize: ymlConfig.UInt("defaultPageSize", 10),
	}

	return &settings
}
