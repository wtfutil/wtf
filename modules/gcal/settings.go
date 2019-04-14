package gcal

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

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

func NewSettingsFromYAML(ymlConfig *config.Config) *Settings {
	localConfig, _ := ymlConfig.Get("wtf.mods.gcal")

	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(ymlConfig),

		conflictIcon:          localConfig.UString("conflictIcon", "🚨"),
		currentIcon:           localConfig.UString("currentIcon", "🔸"),
		displayResponseStatus: localConfig.UBool("displayResponseStatus", true),
		email:                 localConfig.UString("email", ""),
		eventCount:            localConfig.UInt("eventCount", 10),
		multiCalendar:         localConfig.UBool("multiCalendar", false),
		secretFile:            localConfig.UString("secretFile", ""),
		showDeclined:          localConfig.UBool("showDeclined", false),
		textInterval:          localConfig.UInt("textInterval", 30),
		withLocation:          localConfig.UBool("withLocation", true),
	}

	settings.colors.day = localConfig.UString("colors.day", "forestgreen")
	settings.colors.description = localConfig.UString("colors.description", "white")
	settings.colors.past = localConfig.UString("colors.past", "gray")
	settings.colors.title = localConfig.UString("colors.title", "white")

	return &settings
}
