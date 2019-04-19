package cfg

import (
	"github.com/olebedev/config"
)

type Colors struct {
	Background      string
	BorderFocusable string
	BorderFocused   string
	BorderNormal    string
	Checked         string
	HighlightFore   string
	HighlightBack   string
	Text            string
	Title           string
}

type Module struct {
	ConfigKey string
	Name      string
}

type Position struct {
	Height int
	Left   int
	Top    int
	Width  int
}

type Sigils struct {
	CheckedIcon   string
	UncheckedIcon string
}

type Common struct {
	Colors
	Module
	Position
	Sigils

	Enabled         bool
	FocusChar       int
	RefreshInterval int
	Title           string
}

func NewCommonSettingsFromYAML(name, configKey string, ymlConfig *config.Config) *Common {
	colorsPath := "wtf.colors"
	modulePath := "wtf.mods." + configKey
	positionPath := "wtf.mods." + configKey + ".position"
	sigilsPath := "wtf.sigils"

	common := Common{
		Colors: Colors{
			Background:      ymlConfig.UString(modulePath+".colors.background", ymlConfig.UString(colorsPath+".background", "black")),
			BorderFocusable: ymlConfig.UString(colorsPath+".border.focusable", "red"),
			BorderFocused:   ymlConfig.UString(colorsPath+".border.focused", "orange"),
			BorderNormal:    ymlConfig.UString(colorsPath+".border.normal", "gray"),
			Checked:         ymlConfig.UString(colorsPath+".checked", "white"),
			HighlightFore:   ymlConfig.UString(colorsPath+".highlight.fore", "black"),
			HighlightBack:   ymlConfig.UString(colorsPath+".highlight.back", "green"),
			Text:            ymlConfig.UString(modulePath+".colors.text", ymlConfig.UString(colorsPath+".text", "white")),
			Title:           ymlConfig.UString(modulePath+".colors.title", ymlConfig.UString(colorsPath+".title", "white")),
		},

		Module: Module{
			ConfigKey: configKey,
			Name:      name,
		},

		Position: Position{
			Height: ymlConfig.UInt(positionPath + ".height"),
			Left:   ymlConfig.UInt(positionPath + ".left"),
			Top:    ymlConfig.UInt(positionPath + ".top"),
			Width:  ymlConfig.UInt(positionPath + ".width"),
		},

		Sigils: Sigils{
			CheckedIcon:   ymlConfig.UString(sigilsPath+".checkedIcon", "x"),
			UncheckedIcon: ymlConfig.UString(sigilsPath+".uncheckedIcon", " "),
		},

		Enabled:         ymlConfig.UBool(modulePath+".enabled", false),
		FocusChar:       ymlConfig.UInt(modulePath+".focusChar", -1),
		RefreshInterval: ymlConfig.UInt(modulePath+".refreshInterval", 300),
		Title:           ymlConfig.UString(modulePath+".title", name),
	}

	return &common
}
