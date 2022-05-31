package steam

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

const (
	defaultFocusable = true
)

type Settings struct {
	*cfg.Common

	numberOfResults int      `help:"Number of rows to show. Default is 10." optional:"true"`
	key             string   `help:"Steam API key (default is env var STEAM_API_KEY)"`
	userIds         []string `help:"Steam user ids" optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	steam := ymlConfig.UString("steam")
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, steam, defaultFocusable, ymlConfig, globalConfig),

		numberOfResults: ymlConfig.UInt("numberOfResults", 10),
		key:             ymlConfig.UString("key", os.Getenv("STEAM_API_KEY")),
		userIds:         utils.ToStrs(ymlConfig.UList("userIds", make([]interface{}, 0))),
	}
	return &settings
}
