package power

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

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

	str := ""
	str = str + fmt.Sprintf(" %10s: %s\n", "Source", powerSource())
	str = str + "\n"
	str = str + widget.Battery.String()

	widget.View.SetText(fmt.Sprintf("%s", str))
}
