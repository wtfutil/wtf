package zendesk

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Zendesk:

   /: Show/hide this help window
   j: Select the next item in the list
   k: Select the previous item in the list
   o: Open the selected item in a browser

   arrow down: Select the next item in the list
   arrow up:   Select the previous item in the list

   return: Open the selected item in a browser
`

// A Widget represents a Zendesk widget
type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	result   *TicketArray
	selected int
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	ticketArray, err := widget.newTickets()
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
	title := fmt.Sprintf("%s (%d)", widget.CommonSettings.Title, widget.result.Count)
	widget.Redraw(title, widget.textContent(widget.result.Tickets), false)
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
	textColor := widget.settings.common.Colors.Background
	if idx == widget.selected {
		textColor = widget.settings.common.Colors.BorderFocused
	}

	requesterName := widget.parseRequester(ticket)
	str := fmt.Sprintf(" [%s:]%d - %s\n %s\n\n", textColor, ticket.Id, requesterName, ticket.Subject)
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
		ticketURL := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", widget.settings.subdomain, issue.Id)
		wtf.OpenFile(ticketURL)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
}
