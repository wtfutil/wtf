package zendesk

import (
	"fmt"
	"log"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Zendesk ", "zendesk", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */
func (widget *Widget) Refresh() {
	tickets, err := newTickets()
	if err != nil {
		log.Fatal(err)
	}
	widget.UpdateRefreshedAt()

	widget.View.SetTitle(fmt.Sprintf("%s (%d)", widget.Name, len(tickets)))
	widget.View.SetText(widget.textContent(tickets))

}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) textContent(items []string) string {
	if len(items) == 0 {
		return fmt.Sprintf("No unassigned tickets in queue - woop!!")
	}

	str := ""
	for i := range items {
		str = items[i]
	}

	return str
}
