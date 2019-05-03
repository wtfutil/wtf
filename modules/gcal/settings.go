package gcal

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "Calendar"

type colors struct {
	day         string
	description string
	past        string
	title       string

	highlights []interface{}
}

type Settings struct {
	colors
	common *cfg.Common

	conflictIcon          string
	currentIcon           string
	displayResponseStatus bool
	email                 string
	eventCount            int
	multiCalendar         bool
	secretFile            string
	showDeclined          bool
	textInterval          int
	withLocation          bool
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		conflictIcon:          ymlConfig.UString("conflictIcon", "🚨"),
		currentIcon:           ymlConfig.UString("currentIcon", "🔸"),
		displayResponseStatus: ymlConfig.UBool("displayResponseStatus", true),
		email:                 ymlConfig.UString("email", ""),
		eventCount:            ymlConfig.UInt("eventCount", 10),
		multiCalendar:         ymlConfig.UBool("multiCalendar", false),
		secretFile:            ymlConfig.UString("secretFile", ""),
		showDeclined:          ymlConfig.UBool("showDeclined", false),
		textInterval:          ymlConfig.UInt("textInterval", 30),
		withLocation:          ymlConfig.UBool("withLocation", true),
	}

	settings.colors.day = ymlConfig.UString("colors.day", "forestgreen")
	settings.colors.description = ymlConfig.UString("colors.description", "white")
	settings.colors.past = ymlConfig.UString("colors.past", "gray")
	settings.colors.title = ymlConfig.UString("colors.title", "white")

	return &settings
}
