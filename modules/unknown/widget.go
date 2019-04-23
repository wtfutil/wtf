package unknown

import (
	"fmt"

	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	settings *Settings
}

func NewWidget(refreshChan chan<- string, name string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s", widget.Name())))
	widget.View.Clear()

	content := fmt.Sprintf("Widget %s does not exist", widget.Name())
	widget.View.SetText(content)
}
