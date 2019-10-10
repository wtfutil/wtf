package cmdrunner

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "CmdRunner"
)

// Settings for the cmdrunner widget
type Settings struct {
	common *cfg.Common

	args     []string `help:"The arguments to the command, with each item as an element in an array. Example: for curl -I cisco.com, the arguments array would be ['-I', 'cisco.com']."`
	cmd      string   `help:"The terminal command to be run, withouth the arguments. Ie: ping, whoami, curl."`
	tail     bool     `help:"Automatically scroll to the end of the command output."`
	maxLines int      `help:"Maximum number of lines kept in the buffer."`

	// The dimensions of the module
	width  int
	height int
}

// NewSettingsFromYAML loads the cmdrunner portion of the WTF config
func NewSettingsFromYAML(name string, moduleConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, moduleConfig, globalConfig),

		args:     utils.ToStrs(moduleConfig.UList("args")),
		cmd:      moduleConfig.UString("cmd"),
		tail:     moduleConfig.UBool("tail"),
		maxLines: moduleConfig.UInt("maxLines", 256),
	}

	settings.width, settings.height = utils.CalculateDimensions(moduleConfig, globalConfig)

	return &settings
}
