package victorops

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiID  string
	apiKey string
	team   string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.victorops")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiID:  localConfig.UString("apiID", os.Getenv("WTF_VICTOROPS_API_ID")),
		apiKey: localConfig.UString("apiKey", os.Getenv("WTF_VICTOROPS_API_KEY")),
		team:   localConfig.UString("team"),
	}

	return &settings
}
