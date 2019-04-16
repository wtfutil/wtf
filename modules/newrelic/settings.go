package newrelic

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey        string
	applicationID int
	deployCount   int
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.newrelic")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:        localConfig.UString("apiKey", os.Getenv("WTF_NEW_RELIC_API_KEY")),
		applicationID: localConfig.UInt("applicationID"),
		deployCount:   localConfig.UInt("deployCount", 5),
	}

	return &settings
}
