package travisci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "TravisCI"

type Settings struct {
	common *cfg.Common

	apiKey string
	pro    bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey: ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TRAVIS_API_TOKEN"))),
		pro:    ymlConfig.UBool("pro", false),
	}

	return &settings
}
