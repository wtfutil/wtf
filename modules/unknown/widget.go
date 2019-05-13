package unknown

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, pages, settings.common, false),

		settings: settings,
	}

	widget.SetRefreshFunction(widget.Refresh)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	content := fmt.Sprintf("Widget %s and/or type %s does not exist", widget.Name(), widget.CommonSettings.Module.Type)
	widget.Redraw(widget.CommonSettings.Title, content, true)
}
