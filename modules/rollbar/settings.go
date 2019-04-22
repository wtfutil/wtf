package rollbar

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "rollbar"

type Settings struct {
	common *cfg.Common

	accessToken    string
	activeOnly     bool
	assignedToName string
	count          int
	projectName    string
	projectOwner   string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods." + configKey)

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),

		accessToken:    localConfig.UString("accessToken"),
		activeOnly:     localConfig.UBool("activeOnly", false),
		assignedToName: localConfig.UString("assignedToName"),
		count:          localConfig.UInt("count", 10),
		projectName:    localConfig.UString("projectName", "Items"),
		projectOwner:   localConfig.UString("projectOwner"),
	}

	return &settings
}
