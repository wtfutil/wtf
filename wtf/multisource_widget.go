package wtf

import (
	"fmt"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/logger"
)

type MultiSourceWidget struct {
	moduleConfig *cfg.Common
	singular     string
	plural       string

	DisplayFunction func()
	Idx             int
	Sources         []string
}

// NewMultiSourceWidget creates and returns an instance of MultiSourceWidget
func NewMultiSourceWidget(moduleConfig *cfg.Common, singular, plural string) MultiSourceWidget {
	widget := MultiSourceWidget{
		moduleConfig: moduleConfig,
		singular:     singular,
		plural:       plural,
	}

	widget.loadSources()

	logger.Log(fmt.Sprintf("%+v", widget.Sources))

	return widget
}

/* -------------------- Exported Functions -------------------- */

// CurrentSource returns the string representations of the currently-displayed source
func (widget *MultiSourceWidget) CurrentSource() string {
	if widget.Idx >= len(widget.Sources) {
		return ""
	}

	return widget.Sources[widget.Idx]
}

// Next displays the next source in the source list. If the current source is the last
// source it wraps around to the first source
func (widget *MultiSourceWidget) NextSource() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Sources) {
		widget.Idx = 0
	}

	if widget.DisplayFunction != nil {
		widget.DisplayFunction()
	}
}

// Prev displays the previous source in the source list. If the current source is the first
// source, it wraps around to the last source
func (widget *MultiSourceWidget) PrevSource() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Sources) - 1
	}

	if widget.DisplayFunction != nil {
		widget.DisplayFunction()
	}
}

// SetDisplayFunction stores the function that should be called when the source is
// changed. This is typically called from within the initializer for the struct that
// embeds MultiSourceWidget
//
// Example:
//
//   widget := Widget{
//     MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "person", "people")
//   }
//
//   widget.SetDisplayFunction(widget.display)
//
func (widget *MultiSourceWidget) SetDisplayFunction(displayFunc func()) {
	widget.DisplayFunction = displayFunc
}

/* -------------------- Unexported Functions -------------------- */

func (widget *MultiSourceWidget) loadSources() {
	var empty []interface{}

	single := widget.moduleConfig.Config.UString(widget.singular, "")
	multiple := widget.moduleConfig.Config.UList(widget.plural, empty)

	asStrs := ToStrs(multiple)

	if single != "" {
		asStrs = append(asStrs, single)
	}

	widget.Sources = asStrs
}
