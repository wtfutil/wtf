package googleanalytics

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "Google Analytics"

type Settings struct {
	common *cfg.Common

	months         int
	secretFile     string `help:"Your Google client secret JSON file." values:"A string representing a file path to the JSON secret file."`
	viewIds        map[string]interface{}
	enableRealtime bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		months:         ymlConfig.UInt("months"),
		secretFile:     ymlConfig.UString("secretFile"),
		viewIds:        ymlConfig.UMap("viewIds"),
		enableRealtime: ymlConfig.UBool("enableRealtime", false),
	}

	return &settings
}
