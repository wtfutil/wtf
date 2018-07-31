package cryptolive

import (
	"fmt"
	"sync"

	"github.com/senorprogrammer/wtf/cryptoexchanges/cryptolive/price"
	"github.com/senorprogrammer/wtf/cryptoexchanges/cryptolive/toplist"
	"github.com/senorprogrammer/wtf/wtf"
)

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.TextWidget
	priceWidget   *price.Widget
	toplistWidget *toplist.Widget
}

// NewWidget Make new instance of widget
func NewWidget() *Widget {
	widget := Widget{
		TextWidget:    wtf.NewTextWidget(" CryptoLive ", "cryptolive", false),
		priceWidget:   price.NewWidget(),
		toplistWidget: toplist.NewWidget(),
	}

	widget.priceWidget.RefreshInterval = widget.RefreshInterval()
	widget.toplistWidget.RefreshInterval = widget.RefreshInterval()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	var wg sync.WaitGroup

	wg.Add(2)
	widget.priceWidget.Refresh(&wg)
	widget.toplistWidget.Refresh(&wg)
	wg.Wait()

	widget.UpdateRefreshedAt()

	display(widget)
}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	str := ""
	str += widget.priceWidget.Result
	str += widget.toplistWidget.Result
	widget.View.SetText(fmt.Sprintf("\n%s", str))
}
