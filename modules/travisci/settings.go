package travisci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "travisci"

type Settings struct {
	common *cfg.Common

	apiKey string
	pro    bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		apiKey: localConfig.UString("apiKey", os.Getenv("WTF_TRAVIS_API_TOKEN")),
		pro:    localConfig.UBool("pro", false),
	}

	return &settings
}
