package power

import (
	"fmt"

	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	Battery  *Battery
	settings *Settings
}

func NewWidget(refreshChan chan<- string, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(refreshChan, settings.common, false),

		Battery:  NewBattery(),
		settings: settings,
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.Battery.Refresh()

	content := ""
	content = content + fmt.Sprintf(" %10s: %s\n", "Source", powerSource())
	content = content + "\n"
	content = content + widget.Battery.String()

	widget.View.SetText(content)
}
