package ipapi

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const (
	defaultFocusable = false
	defaultTitle     = "IP API"
)

type colors struct {
	name  string
	value string
}

type Settings struct {
	colors
	*cfg.Common
	args []interface{} `help:"Defines what data to display and the order." values:"'ip', 'isp', 'as', 'asName', 'district', 'city', 'region', 'regionName', 'country', 'countryCode', 'continent', 'continentCode', 'coordinates', 'postalCode', 'currency', 'organization', 'timezone' and/or 'reverseDNS'"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		args: ymlConfig.UList("args"),
	}

	settings.colors.name = ymlConfig.UString("colors.name", "red")
	settings.colors.value = ymlConfig.UString("colors.value", "white")
	settings.SetDocumentationPath("ipaddress/ipapi")

	return &settings
}
