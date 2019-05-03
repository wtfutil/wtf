package unknown

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, false),

		app:      app,
		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.app.QueueUpdateDraw(func() {
		widget.View.Clear()

		content := fmt.Sprintf("Widget %s and/or type %s does not exist", widget.Name(), widget.CommonSettings.Module.Type)
		widget.View.SetText(content)
	})
}
