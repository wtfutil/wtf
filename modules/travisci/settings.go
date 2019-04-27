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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey: ymlConfig.UString("apiKey", os.Getenv("WTF_TRAVIS_API_TOKEN")),
		pro:    ymlConfig.UBool("pro", false),
	}

	return &settings
}
