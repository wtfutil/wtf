package github

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	apiKey       string
	baseURL      string
	enableStatus bool
	repositories map[string]interface{}
	uploadURL    string
	username     string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.github")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		apiKey:       localConfig.UString("apiKey", os.Getenv("WTF_GITHUB_TOKEN")),
		baseURL:      localConfig.UString("baseURL", os.Getenv("WTF_GITHUB_BASE_URL")),
		enableStatus: localConfig.UBool("enableStatus", false),
		repositories: localConfig.UMap("repositories"),
		uploadURL:    localConfig.UString("uploadURL", os.Getenv("WTF_GITHUB_UPLOAD_URL")),
		username:     localConfig.UString("username"),
	}

	return &settings
}
