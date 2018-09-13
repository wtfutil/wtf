package wtf

import (
	"fmt"
)

type MultiSourceWidget struct {
	module   string
	singular string
	plural   string

	DisplayFunction func()
	Idx             int
	Sources         []string
}

func NewMultiSourceWidget(module, singular, plural string) MultiSourceWidget {
	return MultiSourceWidget{
		module:   module,
		singular: singular,
		plural:   plural,
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

	s := fmt.Sprintf("wtf.mods.%s.%s", widget.module, widget.singular)
	p := fmt.Sprintf("wtf.mods.%s.%s", widget.module, widget.plural)

	single := Config.UString(s, "")
	multiple := Config.UList(p, empty)

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
