package spotifyweb

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "spotifyweb"

type Settings struct {
	common *cfg.Common

	callbackPort string
	clientID     string
	secretKey    string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		callbackPort: localConfig.UString("callbackPort", "8080"),
		clientID:     localConfig.UString("clientID", os.Getenv("SPOTIFY_ID")),
		secretKey:    localConfig.UString("secretKey", os.Getenv("SPOTIFY_SECRET")),
	}

	return &settings
}
