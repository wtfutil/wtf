package gitter

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Gitter"
)

type Settings struct {
	common *cfg.Common

	apiToken         string `help:"Your Gitter Personal Access Token."`
	numberOfMessages int    `help:"Maximum number of (newest) messages to be displayed. Default is 10" optional:"true"`
	roomURI          string `help:"The room you want to display." values:"Example: wtfutil/Lobby"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiToken:         ymlConfig.UString("apiToken", os.Getenv("WTF_GITTER_API_TOKEN")),
		numberOfMessages: ymlConfig.UInt("numberOfMessages", 10),
		roomURI:          ymlConfig.UString("roomUri", "wtfutil/Lobby"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.apiToken).Load()

	return &settings
}
