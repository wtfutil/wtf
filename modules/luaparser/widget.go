package luaparser

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	widget.View.SetWordWrap(false)

	return widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh redraws the widget content with new data
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	content := "hello cats!"
	content += "\n"

	return widget.CommonSettings().Title, content, true
}
