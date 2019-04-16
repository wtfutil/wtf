package todoist

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey   string
	projects []interface{}
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.todoist")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:   localConfig.UString("apiKey", os.Getenv("WTF_TODOIST_TOKEN")),
		projects: localConfig.UList("projects"),
	}

	return &settings
}
