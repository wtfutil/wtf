package zendesk

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Zendesk widget
type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	result   *TicketArray
	settings *Settings
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	ticketArray, err := widget.newTickets()
	ticketArray.Count = len(ticketArray.Tickets)
	widget.err = err
	widget.result = ticketArray
	widget.Render()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s (%d)", widget.CommonSettings().Title, widget.result.Count)
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	items := widget.result.Tickets
	if len(items) == 0 {
		return title, "No unassigned tickets in queue - woop!!", false
	}

	str := ""
	for idx, data := range items {
		str += widget.format(data, idx)
	}

	return title, str, false
}

func (widget *Widget) format(ticket Ticket, idx int) string {
	textColor := widget.settings.common.Colors.Background
	if idx == widget.GetSelected() {
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

func (widget *Widget) openTicket() {
	sel := widget.GetSelected()
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Tickets) {
		issue := &widget.result.Tickets[sel]
		ticketURL := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", widget.settings.subdomain, issue.Id)
		utils.OpenFile(ticketURL)
	}
}
