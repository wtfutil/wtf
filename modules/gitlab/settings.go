package gitlab

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "gitlab"

type Settings struct {
	common *cfg.Common

	apiKey   string
	domain   string
	projects map[string]interface{}
	username string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", os.Getenv("WTF_GITLAB_TOKEN")),
		domain:   ymlConfig.UString("domain"),
		projects: ymlConfig.UMap("projects"),
		username: ymlConfig.UString("username"),
	}

	return &settings
}
