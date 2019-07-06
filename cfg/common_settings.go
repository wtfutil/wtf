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
	Type string
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
	Position `help:"Defines where in the grid this module’s widget will be displayed."`
	Sigils

	Enabled         bool `help:"Determines whether or not this module is executed and if its data displayed onscreen." values:"true, false"`
	RefreshInterval int  `help:"How often, in seconds, this module will update its data." values:"A positive integer, 0..n." optional:"true"`
	Title           string
	Config          *config.Config

	focusChar int `help:"Define one of the number keys as a short cut key to access the widget." optional:"true"`
}

func NewCommonSettingsFromModule(name, defaultTitle string, moduleConfig *config.Config, globalSettings *config.Config) *Common {
	colorsConfig, _ := globalSettings.Get("wtf.colors")
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
			Type: moduleConfig.UString("type", name),
		},

		Position: NewPositionFromYAML(name, moduleConfig),

		Enabled:         moduleConfig.UBool("enabled", false),
		RefreshInterval: moduleConfig.UInt("refreshInterval", 300),
		Title:           moduleConfig.UString("title", defaultTitle),
		Config:          moduleConfig,

		focusChar: moduleConfig.UInt("focusChar", -1),
	}

	common.Colors.Rows.Even = moduleConfig.UString("colors.rows.even", moduleConfig.UString("rows.even", "white"))
	common.Colors.Rows.Odd = moduleConfig.UString("colors.rows.even", moduleConfig.UString("rows.odd", "lightblue"))

	common.Sigils.Checkbox.Checked = globalSettings.UString(sigilsPath+".checkbox.checked", "x")
	common.Sigils.Checkbox.Unchecked = globalSettings.UString(sigilsPath+".checkbox.unchecked", " ")

	common.Sigils.Paging.Normal = globalSettings.UString(sigilsPath+".paging.normal", globalSettings.UString("wtf.paging.pageSigil", "*"))
	common.Sigils.Paging.Selected = globalSettings.UString(sigilsPath+".paging.select", globalSettings.UString("wtf.paging.selectedSigil", "_"))

	return &common
}

func (common *Common) DefaultFocusedRowColor() string {
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
		sigils += common.Sigils.Paging.Selected
		sigils += strings.Repeat(common.Sigils.Paging.Normal, len-1-pos)

		sigils = "[lightblue]" + fmt.Sprintf(common.RightAlignFormat(width), sigils) + "[white]"
	}

	return sigils
}
