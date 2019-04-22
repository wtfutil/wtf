package cfg

import (
	"fmt"
	"strings"

	"github.com/olebedev/config"
)

type Colors struct {
	Background      string
	BorderFocusable string
	BorderFocused   string
	BorderNormal    string
	Checked         string
	Foreground      string
	HighlightFore   string
	HighlightBack   string
	Text            string
	Title           string

	Rows struct {
		Even string
		Odd  string
	}
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
	Checkbox struct {
		Checked   string
		Unchecked string
	}
	Paging struct {
		Normal   string
		Selected string
	}
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
			Background:      ymlConfig.UString(modulePath+".background", ymlConfig.UString(colorsPath+".background", "black")),
			BorderFocusable: ymlConfig.UString(colorsPath+".border.focusable", "red"),
			BorderFocused:   ymlConfig.UString(colorsPath+".border.focused", "orange"),
			BorderNormal:    ymlConfig.UString(colorsPath+".border.normal", "gray"),
			Checked:         ymlConfig.UString(colorsPath+".checked", "white"),
			Foreground:      ymlConfig.UString(modulePath+".foreground", ymlConfig.UString(colorsPath+".foreground", "white")),
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

		Enabled:         ymlConfig.UBool(modulePath+".enabled", false),
		FocusChar:       ymlConfig.UInt(modulePath+".focusChar", -1),
		RefreshInterval: ymlConfig.UInt(modulePath+".refreshInterval", 300),
		Title:           ymlConfig.UString(modulePath+".title", name),
	}

	common.Colors.Rows.Even = ymlConfig.UString(modulePath+".colors.rows.even", ymlConfig.UString(colorsPath+".rows.even", "white"))
	common.Colors.Rows.Odd = ymlConfig.UString(modulePath+".colors.rows.even", ymlConfig.UString(colorsPath+".rows.odd", "lightblue"))

	common.Sigils.Checkbox.Checked = ymlConfig.UString(sigilsPath+".Checkbox.Checked", "x")
	common.Sigils.Checkbox.Unchecked = ymlConfig.UString(sigilsPath+".Checkbox.Unchecked", " ")

	common.Sigils.Paging.Normal = ymlConfig.UString(sigilsPath+".Paging.Normal", ymlConfig.UString("wtf.paging.pageSigil", "*"))
	common.Sigils.Paging.Selected = ymlConfig.UString(sigilsPath+".Paging.Select", ymlConfig.UString("wtf.paging.selectedSigil", "_"))

	return &common
}

func (common *Common) DefaultFocussedRowColor() string {
	return fmt.Sprintf("%s:%s", common.Colors.HighlightFore, common.Colors.HighlightBack)
}

func (common *Common) DefaultRowColor() string {
	return fmt.Sprintf("%s:%s", common.Colors.Foreground, common.Colors.Background)
}

func (common *Common) RowColor(idx int) string {
	if idx%2 == 0 {
		return common.Colors.Rows.Even
	}

	return common.Colors.Rows.Odd
}

func (common *Common) RightAlignFormat(width int) string {
	return fmt.Sprintf("%%%ds", width-1)
}

func (common *Common) SigilStr(len, pos int, width int) string {
	sigils := ""

	if len > 1 {
		sigils = strings.Repeat(common.Sigils.Paging.Normal, pos)
		sigils = sigils + common.Sigils.Paging.Selected
		sigils = sigils + strings.Repeat(common.Sigils.Paging.Normal, len-1-pos)

		sigils = "[lightblue]" + fmt.Sprintf(common.RightAlignFormat(width), sigils) + "[white]"
	}

	return sigils
}
