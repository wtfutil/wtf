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

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		accessToken:    ymlConfig.UString("accessToken"),
		activeOnly:     ymlConfig.UBool("activeOnly", false),
		assignedToName: ymlConfig.UString("assignedToName"),
		count:          ymlConfig.UInt("count", 10),
		projectName:    ymlConfig.UString("projectName", "Items"),
		projectOwner:   ymlConfig.UString("projectOwner"),
	}

	return &settings
}
