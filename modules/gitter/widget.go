package gitter

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Gitter widget
type Widget struct {
	view.ScrollableWidget

	messages []Message
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.Refresh)
	widget.initializeKeyboardControls()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	room, err := GetRoom(widget.settings.roomURI, widget.settings.apiToken)
	if err != nil {
		widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, err.Error(), true })
		return
	}

	if room == nil {
		widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, "No room", true })
		return
	}

	messages, err := GetMessages(room.ID, widget.settings.numberOfMessages, widget.settings.apiToken)

	if err != nil {
		widget.Redraw(func() (string, string, bool) { return widget.CommonSettings().Title, err.Error(), true })
		return
	}
	widget.messages = messages
	widget.SetItemCount(len(messages))

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("%s - %s", widget.CommonSettings().Title, widget.settings.roomURI)
	if widget.messages == nil || len(widget.messages) == 0 {
		return title, "No Messages To Display", false
	}
	var str string
	for idx, message := range widget.messages {
		row := fmt.Sprintf(
			`[%s] [blue]%s [lightslategray]%s: [%s]%s [aqua]%s`,
			widget.RowColor(idx),
			message.From.DisplayName,
			message.From.Username,
			widget.RowColor(idx),
			message.Text,
			message.Sent.Format("Jan 02, 15:04 MST"),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(message.Text))
	}

	return title, str, true
}
