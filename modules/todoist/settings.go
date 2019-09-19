package todoist

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Todoist"
)

type Settings struct {
	common *cfg.Common

	apiKey   string
	projects []interface{}
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TODOIST_TOKEN"))),
		projects: ymlConfig.UList("projects"),
	}

	return &settings
}
