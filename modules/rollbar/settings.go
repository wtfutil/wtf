package rollbar

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Rollbar"
)

type Settings struct {
	common *cfg.Common

	accessToken    string `help:"Your Rollbar project access token (Only needs read capabilities)."`
	activeOnly     bool   `help:"Only show items that are active." optional:"true"`
	assignedToName string `help:"Set this to your username if you only want to see items assigned to you." optional:"true"`
	count          int    `help:"How many items you want to see. 100 is max." optional:"true"`
	projectName    string `help:"This is used to create a link to the item."`
	projectOwner   string `help:"This is used to create a link to the item."`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		accessToken:    ymlConfig.UString("accessToken"),
		activeOnly:     ymlConfig.UBool("activeOnly", false),
		assignedToName: ymlConfig.UString("assignedToName"),
		count:          ymlConfig.UInt("count", 10),
		projectName:    ymlConfig.UString("projectName", "Items"),
		projectOwner:   ymlConfig.UString("projectOwner"),
	}

	cfg.ModuleSecret(name, globalConfig, &settings.accessToken).Load()

	return &settings
}
