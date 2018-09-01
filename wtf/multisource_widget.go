package wtf

import (
	"fmt"
)

type MultiSourceWidget struct {
	Idx     int
	Sources []string
}

func NewMultiSourceWidget() MultiSourceWidget {
	return MultiSourceWidget{}
}

/* -------------------- Exported Functions -------------------- */

func (widget *MultiSourceWidget) CurrentSource() string {
	return widget.Sources[widget.Idx]
}

func (widget *MultiSourceWidget) LoadSources(module, singular, plural string) {
	var empty []interface{}

	s := fmt.Sprintf("wtf.mods.%s.%s", module, singular)
	p := fmt.Sprintf("wtf.mods.%s.%s", module, plural)

	single := Config.UString(s, "")
	multiple := Config.UList(p, empty)

	asStrs := ToStrs(multiple)

	if single != "" {
		asStrs = append(asStrs, single)
	}

	widget.Sources = asStrs
}
