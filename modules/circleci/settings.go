package circleci

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "CircleCI"
)

type Settings struct {
	common *cfg.Common

	apiKey string `help:"Your CircleCI API token."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey: ymlConfig.UString("apiKey", ymlConfig.UString("apikey", os.Getenv("WTF_CIRCLE_API_KEY"))),
	}

	cfg.ConfigureSecret(
		globalConfig,
		"",
		name,
		nil,
		&settings.apiKey,
	)

	return &settings
}
