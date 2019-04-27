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
	Name string
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
	RefreshInterval int
	Title           string

	focusChar int
}

func NewCommonSettingsFromModule(name string, moduleConfig *config.Config, globalSettings *config.Config) *Common {
	colorsConfig, _ := globalSettings.Get("wtf.colors")
	positionPath := "position"
	sigilsPath := "wtf.sigils"

	common := Common{
		Colors: Colors{
			Background:      moduleConfig.UString("background", globalSettings.UString("background", "black")),
			BorderFocusable: colorsConfig.UString("border.focusable", "red"),
			BorderFocused:   colorsConfig.UString("border.focused", "orange"),
			BorderNormal:    colorsConfig.UString("border.normal", "gray"),
			Checked:         colorsConfig.UString("checked", "white"),
			Foreground:      moduleConfig.UString("foreground", colorsConfig.UString("foreground", "white")),
			HighlightFore:   colorsConfig.UString("highlight.fore", "black"),
			HighlightBack:   colorsConfig.UString("highlight.back", "green"),
			Text:            moduleConfig.UString("colors.text", colorsConfig.UString("text", "white")),
			Title:           moduleConfig.UString("colors.title", colorsConfig.UString("title", "white")),
		},

		Module: Module{
			Name: name,
		},

		Position: Position{
			Height: moduleConfig.UInt(positionPath + ".height"),
			Left:   moduleConfig.UInt(positionPath + ".left"),
			Top:    moduleConfig.UInt(positionPath + ".top"),
			Width:  moduleConfig.UInt(positionPath + ".width"),
		},

		Enabled:         moduleConfig.UBool("enabled", false),
		RefreshInterval: moduleConfig.UInt("refreshInterval", 300),
		Title:           moduleConfig.UString("title", name),

		focusChar: moduleConfig.UInt("focusChar", -1),
	}

	common.Colors.Rows.Even = moduleConfig.UString("colors.rows.even", moduleConfig.UString("rows.even", "white"))
	common.Colors.Rows.Odd = moduleConfig.UString("colors.rows.even", moduleConfig.UString("rows.odd", "lightblue"))

	common.Sigils.Checkbox.Checked = globalSettings.UString(sigilsPath+".Checkbox.Checked", "x")
	common.Sigils.Checkbox.Unchecked = globalSettings.UString(sigilsPath+".Checkbox.Unchecked", " ")

	common.Sigils.Paging.Normal = globalSettings.UString(sigilsPath+".Paging.Normal", globalSettings.UString("wtf.paging.pageSigil", "*"))
	common.Sigils.Paging.Selected = globalSettings.UString(sigilsPath+".Paging.Select", globalSettings.UString("wtf.paging.selectedSigil", "_"))

	return &common
}

func (common *Common) DefaultFocussedRowColor() string {
	return fmt.Sprintf("%s:%s", common.Colors.HighlightFore, common.Colors.HighlightBack)
}

func (common *Common) DefaultRowColor() string {
	return fmt.Sprintf("%s:%s", common.Colors.Foreground, common.Colors.Background)
}

func (common *Common) FocusChar() string {
	focusChar := string('0' + common.focusChar)
	if common.focusChar == -1 {
		focusChar = ""
	}

	return focusChar
}

func (common *Common) RowColor(idx int) string {
	if idx%2 == 0 {
		return common.Colors.Rows.Even
	}

	return common.Colors.Rows.Odd
}

func (common *Common) RightAlignFormat(width int) string {
	borderOffset := 2
	return fmt.Sprintf("%%%ds", width-borderOffset)
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
