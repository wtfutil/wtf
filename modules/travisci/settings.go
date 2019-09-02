package travisci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "TravisCI"

type Settings struct {
	common *cfg.Common

	apiKey  string
	compact bool
	limit   string
	pro     bool
	sort_by string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:  ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_TRAVIS_API_TOKEN"))),
		pro:     ymlConfig.UBool("pro", false),
		compact: ymlConfig.UBool("compact", false),
		limit:   ymlConfig.UString("limit", "10"),
		sort_by: ymlConfig.UString("sort_by", "id:desc"),
	}

	return &settings
}
