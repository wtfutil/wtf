package urlcheck

// Example config
//
// urlcheck:
// 	timeout: 1
// 	urls:
// 		- pli.ski
// 		- www.pli.ski
// 		- http://www.pli.ski
// 		- http://www.nope.nope
// 		- https://httpbin.org/status/500
// 	position:
// 		top: 0
// 		left: 0
// 		height: 2
// 		width: 3
// 	refreshInterval: 40
// 	enabled: true

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "urlcheck"
)

type Settings struct {
	Common *cfg.Common

	requestTimeout int      `help:"Max Request duration in seconds"`
	urls           []string `help:"A list of url to check"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		requestTimeout: ymlConfig.UInt("timeout", 30),
	}
	settings.urls = cfg.ParseAsMapOrList(ymlConfig, "urls")
	return &settings
}
