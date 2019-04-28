package gitter

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "gitter"

type Settings struct {
	common *cfg.Common

	apiToken         string
	numberOfMessages int
	roomURI          string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiToken:         ymlConfig.UString("apiToken", os.Getenv("WTF_GITTER_API_TOKEN")),
		numberOfMessages: ymlConfig.UInt("numberOfMessages", 10),
		roomURI:          ymlConfig.UString("roomUri", "wtfutil/Lobby"),
	}

	return &settings
}
