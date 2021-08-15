package luaparser

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
	lua "github.com/yuin/gopher-lua"
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
	// content := "hello cats!"
	// content += "\n"

	content, err := widget.parse()
	if err != nil {
		content = err.Error()
	}

	return widget.CommonSettings().Title, content, true
}

func (widget *Widget) parse() (string, error) {
	L := lua.NewState()
	defer L.Close()

	output := "this is go"

	L.Push(lua.LString("this is lua"))

	output = L.Get(-1).String()

	return output, nil
}
