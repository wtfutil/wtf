package victorops

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "victorops"

type Settings struct {
	common *cfg.Common

	apiID  string
	apiKey string
	team   string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiID:  ymlConfig.UString("apiID", os.Getenv("WTF_VICTOROPS_API_ID")),
		apiKey: ymlConfig.UString("apiKey", os.Getenv("WTF_VICTOROPS_API_KEY")),
		team:   ymlConfig.UString("team"),
	}

	return &settings
}
