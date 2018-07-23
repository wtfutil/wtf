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

func (widget *Widget) textContent(items []Ticket) string {
	if len(items) == 0 {
		return fmt.Sprintf("No unassigned tickets in queue - woop!!")
	}

	str := ""
	for _, data := range items {
		//str = items[i]
		//str = fmt.Sprintf(data.Id)
		str = str + widget.format(data)
	}

	return str
}

func (widget *Widget) format(ticket Ticket) string {
	var str string
	requesterName := widget.parseRequester(ticket)
	str = fmt.Sprintf(" [green]%d - %s\n %s\n\n", ticket.Id, requesterName, ticket.Subject)
	return str
}

// this is a nasty means of extracting the actual name of the requester from the Via interface of the Ticket.
// very very open to improvements on this
func (widget *Widget) parseRequester(ticket Ticket) interface{} {
	viaMap := ticket.Via
	via := viaMap.(map[string]interface{})
	source := via["source"]
	fromMap, _ := source.(map[string]interface{})
	from := fromMap["from"]
	fromValMap := from.(map[string]interface{})
	fromName := fromValMap["name"]
	return fromName
}
