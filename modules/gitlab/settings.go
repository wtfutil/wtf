package gitlab

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey   string
	domain   string
	projects map[string]interface{}
	username string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.gitlab")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:   localConfig.UString("apiKey", os.Getenv("WTF_GITLAB_TOKEN")),
		domain:   localConfig.UString("domain"),
		projects: localConfig.UMap("projects"),
		username: localConfig.UString("username"),
	}

	return &settings
}
