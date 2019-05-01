package github

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "GitHub"

type Settings struct {
	common *cfg.Common

	apiKey       string
	baseURL      string
	enableStatus bool
	repositories map[string]interface{}
	uploadURL    string
	username     string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:       ymlConfig.UString("apiKey", os.Getenv("WTF_GITHUB_TOKEN")),
		baseURL:      ymlConfig.UString("baseURL", os.Getenv("WTF_GITHUB_BASE_URL")),
		enableStatus: ymlConfig.UBool("enableStatus", false),
		repositories: ymlConfig.UMap("repositories"),
		uploadURL:    ymlConfig.UString("uploadURL", os.Getenv("WTF_GITHUB_UPLOAD_URL")),
		username:     ymlConfig.UString("username"),
	}

	return &settings
}
