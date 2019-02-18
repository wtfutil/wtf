package unknown

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget(app *tview.Application, name string) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, name, name, false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s", widget.Name)))
	widget.View.Clear()

	content := fmt.Sprintf("Widget %s does not exist", widget.Name)
	widget.View.SetText(content)
}
