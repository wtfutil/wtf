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
	*cfg.Common

	args       []string `help:"The arguments to the command, with each item as an element in an array. Example: for curl -I cisco.com, the arguments array would be ['-I', 'cisco.com']."`
	cmd        string   `help:"The terminal command to be run, withouth the arguments. Ie: ping, whoami, curl."`
	tail       bool     `help:"Automatically scroll to the end of the command output."`
	pty        bool     `help:"Run the command in a pseudo-terminal. Some apps will behave differently if they feel in a terminal. For example, some apps will produce colorized output in a terminal, and non-colorized output otherwise. Default false" optional:"true"`
	maxLines   int      `help:"Maximum number of lines kept in the buffer."`
	workingDir string   `help:"Working directory for command to run in" optional:"true"`

	// The dimensions of the module
	width  int
	height int
}

// NewSettingsFromYAML loads the cmdrunner portion of the WTF config
func NewSettingsFromYAML(name string, moduleConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, moduleConfig, globalConfig),

		args:       utils.ToStrs(moduleConfig.UList("args")),
		workingDir: moduleConfig.UString("workingDir", "."),
		cmd:        moduleConfig.UString("cmd"),
		pty:        moduleConfig.UBool("pty", false),
		tail:       moduleConfig.UBool("tail", false),
		maxLines:   moduleConfig.UInt("maxLines", 256),
	}

	width, height, err := utils.CalculateDimensions(moduleConfig, globalConfig)
	if err == nil {
		settings.width = width
		settings.height = height
	}

	return &settings
}
