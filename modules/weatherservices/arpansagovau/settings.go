package arpansagovau

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "ARPANSA UV Data"

type Settings struct {
	common *cfg.Common
	city string;
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),
		city: ymlConfig.UString("locationid"),
	}

	return &settings
}
