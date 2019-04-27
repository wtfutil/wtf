package wtf

import (
	"github.com/wtfutil/wtf/cfg"
)

type MultiSourceWidget struct {
	moduleConfig *cfg.Common
	singular     string
	plural       string

	DisplayFunction func()
	Idx             int
	Sources         []string
}

func NewMultiSourceWidget(moduleConfig *cfg.Common, singular, plural string) MultiSourceWidget {
	return MultiSourceWidget{
		moduleConfig: moduleConfig,
		singular:     singular,
		plural:       plural,
	}
}

/* -------------------- Exported Functions -------------------- */

func (widget *MultiSourceWidget) CurrentSource() string {
	if widget.Idx >= len(widget.Sources) {
		return ""
	}

	return widget.Sources[widget.Idx]
}

func (widget *MultiSourceWidget) LoadSources() {
	var empty []interface{}

	single := widget.moduleConfig.Config.UString(widget.singular, "")
	multiple := widget.moduleConfig.Config.UList(widget.plural, empty)

	asStrs := ToStrs(multiple)

	if single != "" {
		asStrs = append(asStrs, single)
	}

	widget.Sources = asStrs
}

func (widget *MultiSourceWidget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Sources) {
		widget.Idx = 0
	}

	if widget.DisplayFunction != nil {
		widget.DisplayFunction()
	}
}

func (widget *MultiSourceWidget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Sources) - 1
	}

	if widget.DisplayFunction != nil {
		widget.DisplayFunction()
	}
}

func (widget *MultiSourceWidget) SetDisplayFunction(displayFunc func()) {
	widget.DisplayFunction = displayFunc
}
