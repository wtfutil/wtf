package zendesk

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	result   *TicketArray
	selected int
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "Zendesk", "zendesk", true),
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */
func (widget *Widget) Refresh() {
	ticketStatus := wtf.Config.UString("wtf.mods.zendesk.status")
	ticketArray, err := newTickets(ticketStatus)
	ticketArray.Count = len(ticketArray.Tickets)
	if err != nil {
		log.Fatal(err)
	} else {
		widget.result = ticketArray
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.View.SetTitle(fmt.Sprintf("%s (%d)", widget.Name, widget.result.Count))
	widget.View.SetText(widget.textContent(widget.result.Tickets))
}

func (widget *Widget) textContent(items []Ticket) string {
	if len(items) == 0 {
		return fmt.Sprintf("No unassigned tickets in queue - woop!!")
	}

	str := ""
	for idx, data := range items {
		str = str + widget.format(data, idx)
	}

	return str
}

func (widget *Widget) format(ticket Ticket, idx int) string {
	var str string
	requesterName := widget.parseRequester(ticket)
	textColor := wtf.Config.UString("wtf.colors.background", "green")
	if idx == widget.selected {
		textColor = wtf.Config.UString("wtf.colors.background", "orange")
	}
	str = fmt.Sprintf(" [%s:]%d - %s\n %s\n\n", textColor, ticket.Id, requesterName, ticket.Subject)
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

func (widget *Widget) next() {
	widget.selected++
	if widget.result != nil && widget.selected >= len(widget.result.Tickets) {
		widget.selected = 0
	}
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.result != nil {
		widget.selected = len(widget.result.Tickets) - 1
	}
}

func (widget *Widget) openTicket() {
	sel := widget.selected
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Tickets) {
		issue := &widget.result.Tickets[widget.selected]
		ticketUrl := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", subdomain(), issue.Id)
		wtf.OpenFile(ticketUrl)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "j":
		// Select the next item down
		widget.next()
		widget.display()
		return nil
	case "k":
		// Select the next item up
		widget.prev()
		widget.display()
		return nil
	}
	switch event.Key() {
	case tcell.KeyEnter:
		widget.openTicket()
		return nil
	case tcell.KeyEsc:
		// Unselect the current row
		widget.unselect()
		widget.display()
		return event
	default:
		// Pass it along
		return event
	}
}
