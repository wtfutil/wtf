package travisci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	Common *cfg.Common

	apiKey string
	pro    bool
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.travisci")

	settings := Settings{
		Common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey: localConfig.UString("apiKey", os.Getenv("WTF_TRAVIS_API_TOKEN")),
		pro:    localConfig.UBool("pro", false),
	}

	return &settings
}
