package zendesk

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"

	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Zendesk ", "zendesk", true),
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */
func (widget *Widget) Refresh() {
	ticketStatus := wtf.Config.UString("wtf.mods.zendesk.status")
	tickets, ticketArray, err := newTickets(ticketStatus)
	ticketArray.Count = len(ticketArray.Tickets)
	if err != nil {
		log.Fatal(err)
	}
	widget.UpdateRefreshedAt()

	widget.View.SetTitle(fmt.Sprintf("%s (%d)", widget.Name, ticketArray.Count))
	widget.View.SetText(widget.textContent(tickets))

}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) textContent(items []Ticket) string {
	if len(items) == 0 {
		return fmt.Sprintf("No unassigned tickets in queue - woop!!")
	}

	str := ""
	for _, data := range items {
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

func (widget *Widget) openTicket() {
	wtf.OpenFile("https://" + subdomain + ".zendesk.com")
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "f":
		widget.openTicket()
		return nil
	}
	switch event.Key() {
	case tcell.KeyEnter:
		widget.openTicket()
		return nil
	default:
		return event
	}
}
