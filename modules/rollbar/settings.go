package rollbar

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

type Settings struct {
	common *cfg.Common

	accessToken    string
	activeOnly     bool
	assignedToName string
	count          int
	projectName    string
	projectOwner   string
}

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.rollbar")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		accessToken:    localConfig.UString("accessToken"),
		activeOnly:     localConfig.UBool("activeOnly", false),
		assignedToName: localConfig.UString("assignedToName"),
		count:          localConfig.UInt("count", 10),
		projectName:    localConfig.UString("projectName", "Items"),
		projectOwner:   localConfig.UString("projectOwner"),
	}

	return &settings
}
