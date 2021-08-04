package updown

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
	defaultTitle     = "Updown.io"
)

type Settings struct {
	*cfg.Common

	apiKey string   `help:"An Updown API key." optional:"false"`
	tokens []string `help:"Filters the checks and returns only the checks with the specified tokens"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiKey: ymlConfig.UString("apiKey", os.Getenv("WTF_UPDOWN_APIKEY")),
		tokens: utils.ToStrs(ymlConfig.UList("tokens")),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()

	return &settings
}
