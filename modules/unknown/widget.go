package unknown

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	content := fmt.Sprintf("Widget %s and/or type %s does not exist", widget.Name(), widget.CommonSettings().Type)
	widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, content, true })
}
