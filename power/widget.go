package power

import (
	"fmt"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	Battery *Battery
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Power ", "power", false),
		Battery:    NewBattery(),
	}

	widget.View.SetWrap(true)

	return &widget
}

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.Battery.Refresh()

	content := ""
	content = content + fmt.Sprintf(" %10s: %s\n", "Source", powerSource())
	content = content + "\n"
	content = content + widget.Battery.String()

	widget.View.SetText(content)
}
