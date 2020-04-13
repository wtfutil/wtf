package pihole

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = true
	defaultTitle     = "Pi-hole"
)

type Settings struct {
	common         *cfg.Common
	wrapText       bool
	apiUrl         string
	token          string
	showTopItems   int
	showTopClients int
	maxClientWidth int
	maxDomainWidth int
	showSummary    bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common:         cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),
		apiUrl:         ymlConfig.UString("apiUrl"),
		token:          ymlConfig.UString("token"),
		showSummary:    ymlConfig.UBool("showSummary", true),
		showTopItems:   ymlConfig.UInt("showTopItems", 5),
		showTopClients: ymlConfig.UInt("showTopClients", 5),
		maxClientWidth: ymlConfig.UInt("maxClientWidth", 20),
		maxDomainWidth: ymlConfig.UInt("maxDomainWidth", 20),
	}

	return &settings
}
